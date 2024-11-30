export {}; // 將文件轉換為模組

declare global {
    const jQuery: any
    interface Window {
        hljs: any; // 或者更具體的類型
    }
}