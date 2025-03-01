package org.ds;

public class Main {
    public static void main(String[] args) {
        if (args.length != 2) {
            System.out.println("java -jar (jar name) PORT_NUMBER SERVER_NAME");
        }
        int currentServerPort = Integer.parseInt(args[0]);
        String serverName = args[1];

        WebServer webServer = new WebServer(currentServerPort, serverName);

        webServer.startServer();
    }
}