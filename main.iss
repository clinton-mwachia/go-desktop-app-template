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