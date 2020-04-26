import path from 'path'
import { remote } from 'electron'
import { platform } from 'os'


// Get the path that `extraResources` are sent to. This is `<app>/Resources`
// on macOS. remote.app.getAppPath() returns `<app>/Resources/app.asar` so
// we just get the parent directory. If the app is not packaged we just use
// `<current working directory>/resources`.
const resourcesPath = remote.app.isPackaged
  ? path.dirname(remote.app.getAppPath())
  : path.resolve('resources/' + platform())

/*let gitPath = 'git'
if(process.platform == "win32"){
    gitPath = path.join(resourcesPath, 'PortableGit', 'bin', 'git')
}*/

export {resourcesPath}
