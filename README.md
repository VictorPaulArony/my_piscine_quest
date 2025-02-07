# Lem-in Project

This project simulates ant colony movement through a network of interconnected rooms.

## Prerequisites

-   Go installed (version 1.18 or higher)

## Installation

1.  **Clone the repository:**

    ```
    git clone https://learn.zone01kisumu.ke/git/viarony/lem-in.git
    cd lem-in
    ```

## Usage

There are two ways to run this project: using `go run` and using a shell script.

### 1. Using `go run`

1.  **Run the program:**

    ```
    go run main.go <filename>
    ```

    Replace `<filename>` with the path to your input text file (e.g., `text.txt`).

    **Example:**

    ```
    go run main.go text.txt
    ```

### 2. Using the Script

1.  **Change the .txt file content if needed or the file name in the script file  (i.e `text.txt`):**

2.  **Make the script executable:**

    ```
    chmod +x script.sh
    ```

3.  **Run the script:**

    ```
    ./script.sh
    ```

## Input File Format

The input file should define the ant colony's room layout and links as defined below:

### 1. Number of ants

The first line of the file must be a positive number representing the total number of ants in the colony.

### 2. Rooms

Rooms are defined after the number of ants.  Room names must be unique. A room definition looks like:

    `<room_name> <x_coordinate> <y_coordinate>`

-   Start Room: `##start` must come before the room that you want to assign as the start
-   End Room: `##end` must come before the room that you want to assign as the end

### 3. Connections

Connections (links) between rooms are defined after all the rooms have been defined. A connection definition looks like:

    `<room1>-<room2>`

where `<room1>` and `<room2>` are the names of the rooms you are connecting.

### Example text.txt

```text
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1
```

## Output
```bash
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
```
