# CLI Task Manager in Go

A CLI Task Manager is an ideal project to learn Go's practical features, including command-line parsing, file I/O, data persistence, and basic concurrency. Below are comprehensive requirements and suggestions, informed by best practices and real-world examples.

---

## ✅ Core Functional Requirements

### 🔹 Add Task

* **Command:**

  ```bash
  task-cli add "Task description"
  ```
* **Function:** Adds a new task with a unique ID and a default status (e.g., "todo").
* **Output:** Confirmation message with the new task’s ID.

### 🔹 Delete Task

* **Command:**

  ```bash
  task-cli delete <id>
  ```
* **Function:** Removes a task by its unique ID.

### 🔹 Mark Task Status

* **Commands:**

  ```bash
  task-cli mark-status <id> <status>
  ```
* **Function:** Updates the status of a task (e.g., "todo", "in-progress", "done").

### 🔹 List Tasks

* **Commands:**

  ```bash
  task-cli list
  task-cli list done|todo|in-progress
  ```
* **Function:** Lists all or filtered tasks with IDs, descriptions, and statuses.

---

## 📦 Data Model

Each task should have:

* **ID:** UUID (or integer), unique
* **Description:** Required string
* **Status:** Enum or string (`"todo"`, `"in-progress"`, `"done"`)
* **CreatedAt:** Timestamp
* **UpdatedAt:** Timestamp

---

## 💾 Persistence

* **Storage Format:** JSON file (e.g., `tasks.json`)
* **Read/Write:** Load file at startup, write changes after each operation
* **Atomicity:** Ensure atomic writes to avoid corruption

---

## 🧭 CLI Design Suggestions

* **CLI Framework:** Use [`spf13/cobra`](https://github.com/spf13/cobra)
* **Help and Usability:**

  * Support `-help` flag
  * Provide clear error messages
* **Input Handling:**

  * Accept CLI arguments
  * Fallback to interactive prompts if args missing
* **Output:** Use tables for formatting (`go-pretty/table`)
* **Logging:** Optionally log actions to a file
* **Exit Prompts:** Remind users of exit keys (e.g., `Ctrl+C`)

---

## 🛠 Recommended Go Libraries

* [`spf13/cobra`](https://github.com/spf13/cobra) – CLI framework
* `encoding/json` – JSON handling
* `os`, `ioutil`, `bufio` – File and input utilities

---

## 🧱 Development Steps

1. Define `Task` struct with JSON tags
2. Implement file load/save logic
3. Scaffold commands using Cobra
4. Implement logic for each command
5. Add help, usage hints, logging
6. Test for edge cases (duplicates, invalid inputs, missing files)

---

## 📚 References & Inspiration

* [Developer Roadmap Task Tracker](https://roadmap.sh/projects/task-tracker)
* [Gophercises Task Manager](https://github.com/gophercises/task)
* [Harikishan-AI Task Manager](https://github.com/Harikishan-AI/Task-Manager-CLI-Application)
* [YouTube Golang Task Manager](https://www.youtube.com/watch?v=CJgKn3Kebno)
* [Bytesize Go CLI Structure](https://www.bytesizego.com/blog/structure-go-cli-app)
* [Dev.to CLI Article](https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43)
* [Atlassian CLI Design Guide](https://www.atlassian.com/blog/it-teams/10-design-principles-for-delightful-clis)
