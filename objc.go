package main

/*
#cgo darwin CFLAGS: -x objective-c
#cgo darwin LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

void hello() {
	NSLog(@"%@", NSBundle.mainBundle.executableURL.path );
	NSString *appParentDirectory = [[[NSBundle mainBundle] bundlePath] stringByDeletingLastPathComponent];
}
*/
import "C"
import "fmt"

func main() {
	cool := C.hello()
	fmt.Printf("%T %+v %#v %v", cool, cool, cool, cool)
}
