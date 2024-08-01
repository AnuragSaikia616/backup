#include "helpers.h"
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <unistd.h>

void doSomething(int fd) {
  char readerBuf[100];
  size_t n = read(fd, readerBuf, sizeof(readerBuf));
  if (n < 0) {
    perror("read failed");
    exit(EXIT_FAILURE);
  }

  printf("CLIENT: %s\n", readerBuf);
  write(fd, "hi, i am server sending this message", n);
}
int main(int argc, char *argv[]) {
  // Getting the file descriptor using the socket()  sys call
  int fd = socket(AF_INET, SOCK_STREAM, 0);
  if (fd < 0) {
    perror("socker failed");
    exit(EXIT_FAILURE);
  }
  int val = 1;
  // socket configuration
  setsockopt(fd, SOL_SOCKET, SO_REUSEADDR, &val, sizeof(val));

  // define the address and port information in the struct
  struct sockaddr_in addr;
  addr.sin_family = AF_INET;
  addr.sin_port = ntohs(1234);
  addr.sin_addr.s_addr = ntohl(0);

  // bind the file descriptor to the address and port number
  int rv = bind(fd, (const struct sockaddr *)&addr, sizeof(addr));
  if (rv != 0) {
    perror("bind failed");
    exit(EXIT_FAILURE);
  }

  rv = listen(fd, SOMAXCONN);
  if (rv != 0) {
    perror("listen failed");
    exit(EXIT_FAILURE);
  }

  while (1) {
    // define connection address and len
    struct sockaddr_in addr = {};
    socklen_t len = sizeof(addr);

    // get connection file descriptor
    int conn_fd = accept(fd, (struct sockaddr *)&addr, &len);
    if (conn_fd < 0) {
      continue;
    }

    // Handle connections
    while (1) {
      int err = handleRequest(conn_fd);
      if (err != 0) {
        break;
      }
    }

    close(conn_fd);
  }

  return EXIT_SUCCESS;
}
