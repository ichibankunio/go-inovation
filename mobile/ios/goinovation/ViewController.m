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

- (MTKView*)mtkView {
    return (MTKView*)[self.view viewWithTag:101];
}

- (void)viewDidLoad {
    [super viewDidLoad];

    if ([self glkView]) {
        EAGLContext *context = [[EAGLContext alloc] initWithAPI:kEAGLRenderingAPIOpenGLES2];
        [self glkView].context = context;
    
        [EAGLContext setCurrentContext:context];
    
        CADisplayLink *displayLink = [CADisplayLink displayLinkWithTarget:self selector:@selector(drawFrame)];
        [displayLink addToRunLoop:[NSRunLoop currentRunLoop] forMode:NSDefaultRunLoopMode];
    } else if ([self mtkView]) {
        // TODO
    }
}

- (void)viewDidLayoutSubviews {
    [super viewDidLayoutSubviews];

    CGRect viewRect = [[self view] frame];
    double scaleX = viewRect.size.width / (double)MobileScreenWidth;
    double scaleY = viewRect.size.height / (double)MobileScreenHeight;
    double scale = MIN(scaleX, scaleY);
    int width = (int)MobileScreenWidth * scale;
    int height = (int)MobileScreenHeight * scale;
    int x = (viewRect.size.width - width) / 2;
    int y = (viewRect.size.height - height) / 2;
    CGRect kitViewRect = CGRectMake(x, y, width, height);
    [[self glkView] setFrame:kitViewRect];
    [[self mtkView] setFrame:kitViewRect];
    
    if (!MobileIsRunning()) {
        NSError* err = nil;
        MobileStart(scale, &err);
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
    MobileUpdate(&err);
    if (err != nil) {
        NSLog(@"Error: %@", err);
    }
}

- (void)updateTouches:(NSSet*)touches {
    for (UITouch* touch in touches) {
        if (touch.view != [self glkView] && touch.view != [self mtkView]) {
            continue;
        }
        CGPoint location = [touch locationInView:touch.view];
        MobileUpdateTouchesOnIOS(touch.phase, (int64_t)touch, location.x, location.y);
    }
}

- (void)drawInMTKView:(nonnull MTKView *)view {
    NSError* err = nil;
    MobileUpdate(&err);
    if (err != nil) {
        NSLog(@"Error: %@", err);
    }
}

- (void)mtkView:(nonnull MTKView *)view drawableSizeWillChange:(CGSize)size {
    // Do nothing
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
