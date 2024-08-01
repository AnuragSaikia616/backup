import pygame
import math
import random



pygame.init()
pygame.display.set_caption("Partiles")
HEIGHT, WIDTH = 300, 300
win = pygame.display.set_mode((WIDTH, HEIGHT))

class particles:
    def __init__(self,colors,number):
        self.group = []
        self.colors = colors
        self.number = number
        self.create()
    def create(self):
        for i in range(self.number):
            self.group.append([random.randint(0,WIDTH),random.randint(0,HEIGHT),0,0])
    def draw(self,win):
        for coor in self.group:
            pygame.draw.circle(win,self.colors,(coor[0],coor[1]),3)
    def update(self,other,G):
        group = self.group
        other_group = other.group
        for g1 in group:
            fx = fy = 0
            for g2 in other_group:
                dx = g1[0] - g2[0]
                dy = g1[1] - g2[1]
                distance = math.sqrt(dx**2 + dy**2)
                if distance > 0 and distance < 80:
                    force = G / distance
                    fx += force * dx
                    fy += force * dy
            g1[2] = (g1[2] + fx)/2 # x velocity
            g1[3] = (g1[3] + fy)/2 # y velocity
            g1[0] += g1[2] # x position
            g1[1] += g1[3] # y position
            if g1[0] < 0 or g1[0] > WIDTH:
                g1[2] *= -1
            if g1[1] < 0 or g1[1] > HEIGHT:
                g1[3] *= -1


def main():
    clock = pygame.time.Clock()
    run = True
    yellow = particles((255,255,0), 100)
    green = particles((0,255,100),100)
    while run:
        clock.tick(60) 
        win.fill((0,0,0))
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                run = False


        green.update(green, 0.1)
        yellow.update(green, -0.15)
        yellow.update(yellow, -0.1)
        # green.update(yellow, -0.5)
         
        yellow.draw(win)
        green.draw(win)
        pygame.display.update()
    pygame.quit()

main()

