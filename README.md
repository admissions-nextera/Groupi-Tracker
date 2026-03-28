# Groupie Tracker 🎸

A high-performance web application built with **Go** that visualizes data from a RESTful API. This project focuses on data manipulation, client-server synchronization, and efficient backend architecture.

## 🚀 Features

* **Dynamic Data Visualization:** Transforms raw JSON data into an interactive, themed UI (Artists, Locations, Dates, and Relations).
* **Real-Time Search:** Implemented a custom **Client-Server Search Engine**. Unlike basic frontend filters, this triggers a full request-response cycle to the Go backend for data processing.
* **Optimized Performance:** Utilizes an **In-Memory Cache** system to serve data instantly, reducing external API latency and preventing rate-limiting.
* **Resilient Architecture:** Built with robust error handling to ensure the server remains stable even if external API dependencies fail.
* **Responsive UI:** A custom-designed "Dark Gold" interface built with CSS Grid and Flexbox for a premium user experience.



## 🛠️ Tech Stack

* **Backend:** Go (Golang) — utilizing only standard library packages.
* **Frontend:** HTML5, CSS3 (Modern Variables & Animations), Vanilla JavaScript.
* **Data Source:** External RESTful API (Groupie Trackers API).

## 🧠 Technical Highlights

### In-Memory Caching
To meet the requirement of a "non-crashing" website, I implemented a global cache. The application fetches the entire artist dataset once during the server startup. This allows the search functionality to perform complex string matching (names, members, dates) in microseconds without making redundant network calls.

### The Search "Bridge"
The search functionality demonstrates a complete client-server loop:
1.  **Client:** Captures user input via an `input` event listener.
2.  **Request:** Dispatches a `fetch` request to the `/search` endpoint.
3.  **Server:** A Go handler processes the query, performs case-insensitive partial matching, and encodes the results.
4.  **Response:** The client receives JSON and dynamically re-renders the artist grid.

## 🏁 How to Run

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Anasmoner2022/groupie-tracker.git
    ```
2.  **Navigate to the directory:**
    ```bash
    cd groupie-tracker
    ```
3.  **Run the server:**
    ```bash
    go run .
    ```
4.  **Open your browser:**
    Navigate to `http://localhost:8080`
