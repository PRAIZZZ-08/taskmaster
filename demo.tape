Output demo.gif

Set FontSize 14
Set Width 1200
Set Height 600
Set Theme "Dracula"
Set Padding 20
Set BorderRadius 8
Set Margin 20
Set MarginFill "#1a1a2e"
Set Framerate 30
Set PlaybackSpeed 1

# Show available commands
Type "taskmaster"
Enter
Sleep 2s

# Add some tasks
Type "taskmaster add \"Write unit tests for the auth module\""
Enter
Sleep 1s

Type "taskmaster add \"Update the deployment docs\""
Enter
Sleep 1s

Type "taskmaster add \"Fix the login redirect bug\""
Enter
Sleep 1s

# List all tasks
Type "taskmaster list"
Enter
Sleep 2s

# Mark one done
Type "taskmaster done 3"
Enter
Sleep 1s

# List again to show updated state
Type "taskmaster list"
Enter
Sleep 2s

# Delete a task
Type "taskmaster delete 2"
Enter
Sleep 1s

# Final list
Type "taskmaster list"
Enter
Sleep 3s
