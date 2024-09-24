cmake_minimum_required(VERSION 3.10)

project(RandomProject VERSION 1.0)

# Set some variables
set(MY_VAR "Hello, World!")
set(MY_VAR "This will cause a duplicate warning!") # Duplicate variable
set(SOME_PATH "/usr/local/bin")
set(SOME_PATH "/usr/bin") # Duplicate variable

# Set an executable
add_executable(RandomExecutable main.cpp)

# Including directories
include_directories(${SOME_PATH}) # Potential error if SOME_PATH is not correctly defined

# Adding a library
set(MY_LIB "MyLibrary")
add_library(${MY_LIB} src/library.cpp)

# Duplicate definition of MY_LIB
set(MY_LIB "AnotherLibrary")

# Linking the library
target_link_libraries(RandomExecutable PRIVATE ${MY_LIB}) # This may link the last defined MY_LIB
