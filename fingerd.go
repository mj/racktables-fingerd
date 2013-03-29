package main

import(
	"flag"
	"net"
	"bufio"
        "github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"fmt"
)

const(
	user = "finger"
	password = "finger"
	database = "racktables"
	server = "127.0.0.1:3306"
)

func main() {
	flag.Parse();
	
	socket, err := net.Listen("tcp", ":79");
	
	if err != nil {
		panic(err);
	}

	for {
		connection, err := socket.Accept()
		if err != nil {
			continue
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	reader := bufio.NewReader(connection)
	query, _, _ := reader.ReadLine()

	info, err := getContainers(string(query))
	if err != nil {
		connection.Write([]byte(err.Error() + "\n"))
	} else {
		connection.Write([]byte(info))
	}
}

func getContainers(term string) (string, error) {
	query := "SELECT " +
        	"    container.name, " +
		"    fqdn.string_value " +
		"FROM " +
	        "    RackObject vm, " +
	        "    RackObject container, " +
	        "    EntityLink link, " +
		"    AttributeValue fqdn " +
		"WHERE " +
	        "    vm.name LIKE ? " +
	        "    AND vm.id = link.child_entity_id " +
	        "    AND link.child_entity_type = 'object' " +
	        "    AND link.parent_entity_type = 'object' " +
	        "    AND container.id = link.parent_entity_id " +
		"    AND fqdn.attr_id = 3 " +
		"    AND fqdn.object_id = vm.id";

	db := mysql.New("tcp", "", server, user, password, database)

	err := db.Connect()
	if err != nil {
		return "", err
	}

	sel, err := db.Prepare(query)
	if err != nil {
		return "", err
	}

	rows, _, err := sel.Exec("%" + term + "%")

	if err != nil {
		return "", err
	}

	result := ""
	for _, row := range rows {
		container := row[0].([]byte)
		vm := row[1].([]byte);
		
		result = result + fmt.Sprintf("%s: %s\n", vm, container)
	}

	return result, nil
}
