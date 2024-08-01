#include "helpers.h"
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <unistd.h>

int main(int argc, char *argv[]) {

  int fd = socket(AF_INET, SOCK_STREAM, 0);
  if (fd < 0) {
    perror("socket");
    return EXIT_FAILURE;
  }
  struct sockaddr_in addr;
  addr.sin_addr.s_addr = ntohl(0);
  addr.sin_port = ntohs(1234);
  addr.sin_family = AF_INET;

  int rv = connect(fd, (const struct sockaddr *)&addr, sizeof(addr));
  if (rv != 0) {
    perror("connect");
    return EXIT_FAILURE;
  }

  int err = query(fd, "HELLO 1");
  if (err != 0) {
    goto DONE;
  }
  err = query(fd, "HELLO 2");
  if (err != 0) {
    goto DONE;
  }
  err = query(fd, "HELLO 3");
  if (err != 0) {
    goto DONE;
  }

DONE:
  close(fd);

  return EXIT_SUCCESS;
}
