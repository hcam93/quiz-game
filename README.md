# quiz-game

CLI timed quiz game

This simple game works from a command line prompt (only unix based systems, sorry). The point of this was to work with csv files and work with parrallelism in go.

Add quizes in quiz_promblems folder that are two columns:

Question, Answer <- format

There is currently a timer that will quit at 30 seconds. I was able to use

goroutines and channels to quit in the middle of Scanf() from the user.
