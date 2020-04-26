const { spawn } = require("child_process");
import { resourcesPath } from "@/resources"; // Path to resources.ts
import path from "path";
import store from '@/store'
const fkill = require('fkill');
import { launchWindow } from "@/preview"


const hugoBin = path.join(resourcesPath, "hugo");

function newSite(siteDir) {
    //hugo new site /Users/matthieu/Downloads/donnasummer -v --log --logFile=logs --force
    let hugoParams = [
        "new",
        "site",
        siteDir,
        "-v",
        "--log",
        "--force",
    ];
    try {
        spawn(hugoBin, hugoParams);
    } catch (e) {
        console.error('hugo new site error: ', e)
    }
}

function runserver(siteDir) {
    store.commit("toggleServerIsLaunching", true)
    let hugoParams = [
        "server",
        "--navigateToChanged",
        "-s",
        siteDir,
        "-vw",
        "--cleanDestinationDir",
        "--disableFastRender",
        "--bind=localhost",
        "-p",
        "1313"
    ];

    console.log(hugoBin, hugoParams.join(' '));

    // this.stdout.push(hugoParams);

    let hugoProcess = spawn(hugoBin, hugoParams, { cwd: siteDir });
    console.log(hugoProcess)

    // this.setHugoServer(process);
    console.info("spawned process", hugoProcess.pid);
    //
    hugoProcess.stdout.on("data", (data) => {
        console.log(String(data))
        store.commit('logToConsole', String(data))
        if (String(data).match(/spawned successfully/)) {
            console.log('process launched successfully')
        }
        var rx = /localhost:(.*)\//g;
        var arr = rx.exec(String(data));
        if (arr) {
            console.log(arr[1])
            store.commit('setHugoServerPort', parseInt(arr[1]))
            store.commit('setHugoServer', hugoProcess.pid)

            store.commit("toggleServerIsLaunching", false)
            launchWindow('localhost', parseInt(arr[1]))
        }
    });
    //
    hugoProcess.stderr.on("data", (data) => {
        console.log(String(data))
        store.commit('logToConsole', String(data))
    });
    //
    hugoProcess.on("close", (code) => {
        console.info(`child process closed with code ${code}`);
        store.commit('logToConsole', `child process closed with code ${code}`)
        store.commit('setHugoServerPort', null)
        store.commit('setHugoServer', null)
        store.commit("toggleServerIsLaunching", false)
    });
    hugoProcess.on("error", (code) => {
        console.error(`child process errored with code ${code}`);
        store.commit('logToConsole', `child process errored with code ${code}`)
        store.commit('setHugoServerPort', null)
        store.commit('setHugoServer', null)
        store.commit("toggleServerIsLaunching", false)
    });



}
function stopserver() {
    console.log("hugostop");

    // this.errors = [];
    // this.stdout = [];

    /*this.hugoServer.kill('SIGINT')*/
    try {
        (async () => {
            await fkill([store.state.hugoServerPID], { force: true });
            // await fkill([store.state.hugoServerPID, `:${store.state.hugoServerPID}`],{force: true});
            console.log('Killed process');
        })();
        console.log("Killed process", store.state.hugoServerPID);
        store.commit('setHugoServer', null);
        store.commit('setHugoServerPort', null);
        if (store.state.subwindow !== null) {
            store.state.subwindow.close();
            store.commit('setSubWindow', null)
        }
    } catch (e) {
        console.error('could not kill hugo server', e)
    }


}
export { newSite, runserver, stopserver }

