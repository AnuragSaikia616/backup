#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

const size_t MAX_MSG_SIZE = 4096;

static int read_all(int fd, char *buf, int n) {
  int rv = 0;
  while (n > 0) {
    rv = read(fd, buf, n);
    if (rv <= 0) {
      return -1;
    }
    n -= rv;
    buf += rv;
  }
  return 0;
}

static int write_all(int fd, char *buf, int size) {
  while (size > 0) {
    int rv = write(fd, buf, size);
    if (rv <= 0) {
      return -1;
    }
    size -= rv;
    buf += rv;
  }
  return 0;
}

static int handleRequest(int fd) {
  char rbuf[4 + MAX_MSG_SIZE + 1];
  // Read header
  int err = read_all(fd, rbuf, 4);
  if (err != 0)
    return err;

  // Check if the message size is too big
  uint32_t len;
  memcpy(&len, rbuf, 4);
  if (len > MAX_MSG_SIZE) {
    printf("too big\n");
    return -1;
  }

  err = read_all(fd, rbuf + 4, len);
  if (err != 0) {
    return err;
  }
  printf("SERVER: %s\n", rbuf + 4);

  const char text[] = "Hi from the server";
  len = sizeof(text);
  char wbuf[4 + len];
  memcpy(wbuf, &len, 4);
  memcpy(wbuf + 4, text, len);
  err = write_all(fd, wbuf, sizeof(wbuf));
  return err;
}

static int query(int fd, const char *text) {
  uint32_t len = (uint32_t)sizeof(text);
  if (len > MAX_MSG_SIZE) {
    printf("input too big\n");
    return -1;
  }

  char writeBuf[4 + MAX_MSG_SIZE];
  // copthe the header and then the text
  memcpy(writeBuf, &len, 4);
  memcpy(writeBuf + 4, text, len);
  if (write_all(fd, writeBuf, 4 + len) < 0) {
    return -1;
  }

  char readBuf[4 + MAX_MSG_SIZE + 1];

  // Read header
  int err = read_all(fd, readBuf, 4);
  if (err != 0)
    return err;

  // Check if the message size is too big
  memcpy(&len, readBuf, 4);
  if (len > MAX_MSG_SIZE) {
    printf("too big\n");
    return -1;
  }

  err = read_all(fd, readBuf + 4, len);
  if (err != 0) {
    return err;
  }
  printf("SERVER: %s\n", readBuf + 4);
  return 0;
}
