# Anime Wallpaper Changer

A simple command-line application to automatically change your desktop wallpaper to a random anime wallpaper every time you log in.

## Features

-   Automatically changes your desktop wallpaper on login.
-   Fetches high-quality anime wallpapers from [Wallhaven](https://wallhaven.cc).
-   Interactive first-time setup to configure your preferences.
-   Cross-platform support for Windows, macOS, and Linux.
-   Option to automatically run on login.

## Installation

1.  **Prerequisites:** Make sure you have [Go](https://golang.org/dl/) installed on your system.

2.  **Build the application:** Open your terminal or command prompt, navigate to the project directory, and run the following command:

    ```
    go build
    ```

    This will create an executable file named `anime-wallpaper-changer` (or `anime-wallpaper-changer.exe` on Windows) in the project directory.

## Usage

1.  **Run the application:** Run the executable from your terminal:

    ```
    ./anime-wallpaper-changer
    ```

2.  **First-Time Setup:** The first time you run the application, it will guide you through an interactive setup process:

    *   **Choose your favorite categories:** Select the wallpaper categories you want to see (e.g., General, Anime, People).
    *   **Choose the purity level:** Select the purity level of the wallpapers (e.g., SFW, Sketchy).
    *   **Run on login:** Choose whether you want the application to run automatically every time you log in.

3.  **Enjoy:** The application will now run in the background and change your wallpaper every time you log in based on your preferences.

## Supported Platforms

-   Windows
-   macOS
-   Linux (GNOME)

## Contributing

Contributions are welcome! If you want to add support for other desktop environments on Linux or improve the application in any other way, feel free to open a pull request.
