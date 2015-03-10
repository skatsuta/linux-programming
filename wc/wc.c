#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	int i;

	for (i = 1; i < argc; i++) {
		FILE *f;
		int c;
		int cnt = 0;

		f = fopen(argv[i], "r");
		if (!f) {
			perror(argv[i]);
			exit(1);
		}

		while ((c = fgetc(f)) != EOF) {
			if (c == '\n') {
				cnt++;
			}
		}

		printf("%s: %d lines\n", argv[i], cnt);

		fclose(f);
	}

	exit(0);
}
