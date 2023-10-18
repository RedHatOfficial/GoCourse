    hello = dlsym(library, "hello");

    if (hello != NULL) {
        printf("address for 'hello' retrieved: %p\n", (void*)hello);
        puts("Calling 'hello'...");
        hello();
        puts("...called");
    } else {
        puts("unable to retrieve address for 'hello'");
    }


    if (library != NULL) {
        int err = dlclose(library);
        if (err != 0) {
            puts("unable to close dynamic library");
            return 1;
        } else {
            puts("dynamic library closed");
        }
    }

    return EXIT_SUCCESS;
}
