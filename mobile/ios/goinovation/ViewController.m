//
//  ViewController.m
//  goinovation
//
//  Created by Hajime Hoshi on 6/16/16.
//  Copyright © 2016 Hajime Hoshi. All rights reserved.
//

#import "ViewController.h"

#import "Mobile/Mobile.h"

@interface ViewController ()

@end

@implementation ViewController

- (GLKView*)glkView {
    return (GLKView*)[self.view viewWithTag:100];
}

- (void)viewDidLoad {
    [super viewDidLoad];

    EAGLContext *context = [[EAGLContext alloc] initWithAPI:kEAGLRenderingAPIOpenGLES2];
    [self glkView].context = context;
    
    [EAGLContext setCurrentContext:context];
    
    CADisplayLink *displayLink = [CADisplayLink displayLinkWithTarget:self selector:@selector(drawFrame)];
    [displayLink addToRunLoop:[NSRunLoop currentRunLoop] forMode:NSDefaultRunLoopMode];
}

- (void)viewDidLayoutSubviews {
    [super viewDidLayoutSubviews];
    CGRect viewRect = [[self view] frame];
    double scaleX = (double)viewRect.size.width / (double)GoMobileScreenWidth;
    double scaleY = (double)viewRect.size.height / (double)GoMobileScreenHeight;
    double scale = MAX(1, MIN(scaleX, scaleY));
    int width = (int)GoMobileScreenWidth * scale;
    int height = (int)GoMobileScreenHeight * scale;
    int x = (viewRect.size.width - width) / 2;
    int y = (viewRect.size.height - height) / 2;
    CGRect glkViewRect = CGRectMake(x, y, width, height);
    [[self glkView] setFrame:glkViewRect];
    
    if (!GoMobileIsRunning()) {
        NSError* err = nil;
        GoMobileStart(scale, &err);
        if (err != nil) {
            NSLog(@"Error: %@", err);
        }
    }
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

- (void)drawFrame{
    [[self glkView] setNeedsDisplay];
}

- (void)glkView:(GLKView *)view drawInRect:(CGRect)rect {
    NSError* err = nil;
    GoMobileUpdate(&err);
    if (err != nil) {
        NSLog(@"Error: %@", err);
    }
}

- (void)updateTouches:(NSSet*)touches {
    for (UITouch* touch in touches) {
        if (touch.view != [self glkView]) {
            continue;
        }
        CGPoint location = [touch locationInView:[self glkView]];
        GoMobileUpdateTouchesOnIOS(touch.phase, (int64_t)touch, location.x, location.y);
    }
}

- (void)touchesBegan:(NSSet*)touches withEvent:(UIEvent*)event {
    [self updateTouches:touches];
}

- (void)touchesMoved:(NSSet*)touches withEvent:(UIEvent*)event {
    [self updateTouches:touches];
}

- (void)touchesEnded:(NSSet*)touches withEvent:(UIEvent*)event {
    [self updateTouches:touches];
}

- (void)touchesCancelled:(NSSet*)touches withEvent:(UIEvent*)event {
    [self updateTouches:touches];
}

@end
