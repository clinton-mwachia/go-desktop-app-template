# Go Desktop App Template

A robust Go desktop application developed using the Fyne framework. This app provides various functionalities and features that enhance user experience, such as user management, task management, notifications, settings, and more.

## Features

- **User Management**: Create, update, delete, and manage user roles.
- **Task Management**: Add, edit, delete, and filter tasks.
- **Notifications**: Real-time notifications for user actions, stored in the database and displayed with a counter.
- **Data Export**: Export data to CSV, JSON, etc.
- **Search, Filter, and Sort**: Search, filter, and sort data dynamically.
- **Settings**: Customizable settings for theme, font, etc.
- **User Profile**: Update user details and change passwords.
- **Secure Authentication**: Password hashing and validation using bcrypt.
- **Responsive UI**: Adaptive design for different screen sizes.
- **Accessibility**: Customizable themes and fonts for better accessibility.
- **Sound Alerts**: Audio alerts for notifications and critical actions.

## Prerequisites

To build and run this application, ensure you have the following installed:

- **Go** (v1.16 or later)
- **Fyne** (v2.x)

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/clinton-mwachia/go-desktop-app-template.git
cd go-desktop-app-template
```

### Install Dependencies

```bash
go get fyne.io/fyne/v2
go get golang.org/x/crypto/bcrypt
go get go.mongodb.org/mongo-driver/mongo
```

### Build the Application

```bash
go build -o GoDesktopApp
```

## Packaging the Application

```bash
fyne package -os windows
```
### Additional Packaging Options
The `fyne` package command can be enhanced with various flags for more customization:
1. `-icon Icon.png`: Specify a different icon file.
2. `-appID <id>`: Override the app ID from FyneApp.toml.
3. `-name <name>`: Override the app name.
4. `-release`: Build a release version with optimizations.
5. `-executable <file>`: Specify the executable name if different from the default.

### Packaging for Different Platforms

#### **Windows (.exe Installer)**

1. **Download and Install Inno Setup**: [jrsoftware.org](https://jrsoftware.org/isinfo.php)
2. **Create a New Script**: Use Inno Setup to create a new installer script.
3. **Add Application Files**: Include your compiled `.exe` and other required files.
4. **Compile the Installer**: Follow the wizard to generate an `.exe` installer.

Example Inno Setup Script:
```ini
[Setup]
AppName=Go Desktop App
AppVersion=1.0
DefaultDirName={pf}\desktop-app-template
DefaultGroupName=desktop-app-template
OutputDir=output
OutputBaseFilename=desktop-app-templateSetup
Compression=lzma
SolidCompression=yes

[Files]
Source: "desktop-app-template.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "assets\*"; DestDir: "{app}\assets"; Flags: ignoreversion recursesubdirs createallsubdirs

[Icons]
Name: "{group}\Go Desktop App"; Filename: "{app}\desktop-app-template.exe"
Name: "{group}\Uninstall Go Desktop App"; Filename: "{uninstallexe}"
```

#### **macOS (.dmg Installer)**

1. **Install `appdmg`**:
   ```bash
   npm install -g appdmg
   ```

2. **Create Configuration File (`appdmg.json`)**:
   ```json
   {
     "title": "Go Desktop App",
     "icon": "assets/icon.icns",
     "background": "assets/background.png",
     "icon-size": 80,
     "contents": [
       { "x": 448, "y": 344, "type": "link", "path": "/Applications" },
       { "x": 192, "y": 344, "type": "file", "path": "GoDesktopApp.app" }
     ]
   }
   ```

3. **Build the `.dmg`**:
   ```bash
   appdmg appdmg.json GoDesktopApp.dmg
   ```

#### **Linux (.deb or .rpm Installer)**

**For `.deb` (Debian/Ubuntu):**

1. **Set Up Directory Structure**:
   ```bash
   mkdir -p GoDesktopApp/DEBIAN
   mkdir -p GoDesktopApp/usr/local/bin
   mkdir -p GoDesktopApp/usr/share/applications
   ```

2. **Create Control File**:
   ```plaintext
   Package: godesktopapp
   Version: 1.0
   Section: base
   Priority: optional
   Architecture: amd64
   Depends: libappindicator3-1, libnotify-bin
   Maintainer: Your Name <youremail@example.com>
   Description: A robust Go desktop application
   ```

3. **Build the `.deb`**:
   ```bash
   dpkg-deb --build GoDesktopApp
   ```

## Making the Application Installable

Once you have packaged the application for your target platform, distribute the installer file (`.exe`, `.dmg`, `.deb`, etc.) to users. They can then run the installer to install your Go desktop application on their system.

## Running the Application

After installation, you can run the application from the system's application menu or by executing the installed binary directly.

```bash
./GoDesktopApp   # Linux/macOS
GoDesktopApp.exe # Windows
```

## License

This project is licensed under the Appache-2.0 License - see the [LICENSE](LICENSE) file for details.

