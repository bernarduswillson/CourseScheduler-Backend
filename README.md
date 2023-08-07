# Course Scheduler🎲
>Tugas Seleksi IRK
## Table of Contents
* [Contributors](#contributors)
* [General Information](#general-information)
* [Local Setup](#local-setup)
* [Dynamic Programming](#dynamic-programming)
* [Algorithm](#algorithm)
* [Frameworks and Technologies](#frameworks-and-technologies)
* [References](#references)
## Contributors
| NIM | Nama |
| :---: | :---: |
| 13521021 | Bernardus Willson  |
## General Information 
The Course Scheduler is a program designed to assist students in planning their academic course schedules optimally. This program takes input data on courses along with important information such as the number of credit hours (SKS), the predicted level of difficulty, and the minimum semester required to take the course. Utilizing Dynamic Programming techniques, the program can calculate and determine the best combination of courses that a student can take in one semester, given a specific SKS constraint.

The main objective of this program is to find the combination of courses that yields the highest predicted score per SKS, enabling students to achieve the optimal grade while staying within the maximum SKS limit. The program's output will display the recommended courses along with the total predicted score per SKS obtained from the selected combination.

By using the Course Scheduler, students can efficiently plan their course schedules and maximize their academic achievements.

## Local Setup
<br>
1. Clone BE repo using the command below: 

```
git clone https://github.com/bernarduswillson/CourseScheduler-Backend.git
```
<br>
2. Make sure you have docker installed, run the backend server using the command below:

```
docker-compose up -d
```
or alternatively, you can run the backend server manually using the command below:

```
go run main.go
```
<br>
3. The project will be served on port 8080

```
http://localhost:8080/
```

## Dynamic Programming

Dynamic programming is a powerful optimization technique used to solve complex problems by breaking them down into smaller overlapping subproblems and then storing their solutions in a table for efficient retrieval. It is particularly useful for problems that exhibit overlapping subproblems and optimal substructure.

The key idea behind dynamic programming is to avoid redundant calculations by storing the solutions to subproblems in a data structure (usually an array or a table) and reusing them whenever needed. This approach significantly reduces the time complexity of solving the problem.

## Algorithm
1. Initialize the DP table: The algorithm begins by initializing a 2D slice dp to store the maximum score achievable with a specific number of SKS (credit hours) for the first i courses. The dimensions of the DP table are n+1 (number of courses + 1) and maxSKS+1 (maximum number of SKS allowed + 1).

2. Fill the DP table: The algorithm fills the DP table using bottom-up dynamic programming. It iterates through the courses and the SKS to compute the maximum score that can be achieved with a certain number of SKS. For each course, it considers whether to include it in the selection based on the maximum score achieved without the course and the maximum score achieved by including the course and reducing the available SKS accordingly.

3. Backtrack to find the selected courses: After filling the DP table, the algorithm backtracks to find the courses that contribute to the maximum score. It starts from the bottom-right corner of the DP table (corresponding to the maximum number of SKS) and traces back to the top-left corner (corresponding to the first course). During the backtracking process, it adds the courses that were included in the optimal selection to the selectedCourses list.

4. Calculate total score: The algorithm then calls the totalScore function, which calculates the total score of the selected courses. The total score is calculated as the sum of the products of each course's SKS and its corresponding prediksi score.

5. Helper functions: The algorithm also includes helper functions mapPrediksiToScore and max. mapPrediksiToScore maps a prediksi (grade prediction) string to its corresponding score. max is a simple helper function to find the maximum of two floats.

## Frameworks and Technologies
* [Next.js](https://nextjs.org/)
* [React.js](https://reactjs.org/)
* [Golang](https://golang.org/)
* [Gorm](https://gorm.io/)
* [Gin](https://gin-gonic.com/)
* [MySQL](https://www.mysql.com/)
* [Docker](https://www.docker.com/)

## References
* [0/1 Knapsack Problem Dynamic Programming](https://youtu.be/8LusJS5-AGo)