package cn.kamakura.aservice;

import cn.kamakura.aservice.handler.GreetHandler;
import cn.kamakura.aservice.handler.HelloHandler;
import com.sun.net.httpserver.HttpServer;

import java.io.IOException;
import java.net.InetSocketAddress;

public class Application {

    private static final int DEFAULT_PORT = 8080;

    public static void main(String[] args) throws IOException {
        int port = getPort();
        HttpServer server = HttpServer.create(new InetSocketAddress(port), 0);

        server.createContext("/api/hello", new HelloHandler());
        server.createContext("/api/greet", new GreetHandler());

        server.setExecutor(null);
        server.start();

        System.out.println("Server started on port " + port);
        System.out.println("API endpoints:");
        System.out.println("  GET http://localhost:" + port + "/api/hello");
        System.out.println("  GET http://localhost:" + port + "/api/greet?name=xxx");
    }

    private static int getPort() {
        String portEnv = System.getenv("PORT");
        if (portEnv != null && !portEnv.isEmpty()) {
            try {
                return Integer.parseInt(portEnv);
            } catch (NumberFormatException e) {
                // ignore
            }
        }
        return DEFAULT_PORT;
    }
}
