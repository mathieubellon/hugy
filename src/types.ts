// Store

export interface RootState {
    serverIsLaunching: Boolean,
    currentDir: string | null
    projectPath: string
    currentProject?: Project | null,
    currentProjectFiles?: File[] | null,
    readDirPath?: string | null,
    currentFile?: File | null,
    hugoServerPID?: Number | null,
    hugoServerPort?: Number | null,
    subwindow?: any | null,
    console: any[],
    user: User | null,
}

// Models
export interface File {
    path: string
}

export interface User {
    email?: string
    token?: string
}

export interface Project {
    uuid: string
    name: string
    remoteRepo: string
    localRepo: string
    exist: boolean
    isCloning: boolean
}

