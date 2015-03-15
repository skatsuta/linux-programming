#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netdb.h>
#include <string.h>

static int open_connection(const char *host, char *service);

int main(int argc, char const* argv[])
{
	const char *host;
	int sock;
	FILE *f;
	char buf[1024];

	host = argc > 1 ? argv[1] : "localhost";
	
	if ((sock = open_connection(host, "daytime")) < 0) {
		fprintf(stderr, "socket(2)/connect(2) failed\n");
		exit(1);
	}

	f = fdopen(sock, "r");
	if (!f) {
		perror("fdopen(3)");
		exit(1);
	}

	fgets(buf, sizeof buf, f);
	fclose(f);
	fputs(buf, stdout);
	return 0;
}

static int open_connection(const char *host, char *service) {
	int sock;
	struct addrinfo hints, *res, *ai;
	int err;

	memset(&hints, 0, sizeof(struct addrinfo));
	hints.ai_family = AF_UNSPEC;
	hints.ai_socktype = SOCK_STREAM;

	if ((err = getaddrinfo(host, service, &hints, &res)) != 0) {
		fprintf(stderr, "getaddrinfo(3): %s\n", gai_strerror(err));
		exit(1);
	}

	for (ai = res; ai; ai = ai->ai_next) {
		sock = socket(ai->ai_family, ai->ai_socktype, ai->ai_protocol);
		if (sock < 0) {
			continue;
		}

		if (connect(sock, ai->ai_addr, ai->ai_addrlen) < 0) {
			close(sock);
			continue;
		}

		// success
		freeaddrinfo(res);
		return sock;
	}

	freeaddrinfo(res);
	return -1;
}

