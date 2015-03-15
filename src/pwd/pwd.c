#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>

#define INIT_BUFSIZE 1024

static char* my_getcmd(void);

int main(int argc, char const* argv[])
{
	char *buf = my_getcmd();
	if (!buf) {
		perror("error");
	}

	printf("%s\n", buf);
	return 0;
}

static char* my_getcmd(void) {
	char *buf, *tmp;
	size_t size = INIT_BUFSIZE;

	buf = malloc(size);
	if (!buf) {
		return NULL;
	}

	for (;;) {
		errno = 0;

		if (getcwd(buf, size)) { return buf; }

		if (errno != ERANGE) { break; }

		// re-allocate buffer to double size
		size *= 2;
		tmp = realloc(buf, size);
		if (!tmp) { break; }
		buf = tmp;
	}

	free(buf);
	return NULL;
}

