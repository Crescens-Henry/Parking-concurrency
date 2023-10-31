# Concurrent Parking Lot Simulation
![image](https://github.com/Crescens-Henry/Parking-concurrency/assets/94093619/99fffc96-cce6-483e-a765-21380e84dc11)

This Go project simulates a parking lot using concurrent goroutines. It demonstrates various concurrent programming concepts, including the use of goroutines, channels, and mutex locks. The simulation allows cars to enter, park, and exit a parking lot with a limited number of spaces. It also visualizes the process using the Fyne graphical user interface (GUI) library.

## Features
Goroutines: The project uses goroutines to simulate the concurrent behavior of cars entering and exiting the parking lot. Each car is represented as a goroutine, which can enter, park, and exit the lot independently.

Channels: Channels are utilized for communication between cars and the parking lot. The parking lot restricts the number of available spaces using a channel, and cars signal their entry and exit through channels as well.

Mutex Locks: A mutex lock is used to protect shared resources, ensuring that only one car can access certain parts of the code at a time. This prevents race conditions and data corruption.

## Usage
To run the simulation, follow these steps:

Make sure you have Go installed on your system.

Clone this repository to your local machine.

Build the project:

bash
Copy code:
```
go build
```
Run the executable:
The program will simulate cars entering and exiting the parking lot. The graphical interface displays the parking lot's state in real-time.

## Acknowledgments
This project was developed as part of a concurrent programming course. Special thanks to the Go community and the Fyne developers for providing tools and libraries to make this project possible.

