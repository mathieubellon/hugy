const electron = require("electron");
const remote = electron.remote;

export function launchWindow(host, port) {

    var win = remote.getCurrentWindow();
    let previewWindow = new remote.BrowserWindow({
        width: 800,
        height: 800
    });

    // set subwindow in store

    previewWindow.show();
    previewWindow.on("closed", () => {
        console.log("window closed");
        //this.setSubWindow(null);
    });

    // previewWindow.setBounds({
    //     y: win.getBounds().y,
    //     x: win.getBounds().x + win.getBounds().width
    // });
    // We setTimeout to wait Hugo Server to spin up correctly
    previewWindow.focus();
    console.log("childwindow loading url " + `http://${host}:${port}`);
    previewWindow.loadURL(`http://${host}:${port}`);

}
