import pygame
import math

pygame.init()
WIDTH, HEIGHT = 2000,2000
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("Planet Simulation")
FONT = pygame.font.SysFont("comicsans", 16)


AU = 149.6e6 * 1000
G = 6.67428e-11 
SCALE = 150 / AU
TIMESTEP = 60 * 60 * 24

class planet:
    def __init__(self,name,x,y,radius,mass,color,yvel,sun=False):
        self.name = name
        self.x = x
        self.y = y
        self.radius = radius
        self.mass = mass
        self.color = color
        self.sun = sun
        self.xvel = G * TIMESTEP
        self.yvel = yvel * 1
        self.orbit = []
        self.distance_from_sun = x

    def draw(self,win):
        x = self.x * SCALE + WIDTH / 2 
        y = self.y * SCALE + HEIGHT / 2
        name_text = FONT.render(self.name,1,(255,255,255))
        dis_text = FONT.render(str(round(self.distance_from_sun/1000000,2))+" km",1,(255,255,255))

        pygame.draw.circle(win,self.color,(x,y),self.radius)
        if self.sun == False:
            win.blit(dis_text,(x,y))

        if len(self.orbit) > 2:
            pygame.draw.lines(win, self.color, False, self.orbit, 2)


    def attraction(self,other):
        other_x, other_y = other.x, other.y
        distance_x = other_x - self.x
        distance_y = other_y - self.y
        distance = math.sqrt(distance_x ** 2 + distance_y ** 2)
        

        force = G * self.mass * other.mass / distance**2
        theta = math.atan2(distance_y, distance_x)
        force_x = force * math.cos(theta)
        force_y = force * math.sin(theta)
        return force_x, force_y
    
    def update_position(self,planets):
        total_fx = total_fy = 0
        for planet in planets:
            if self == planet:
                continue
            fx, fy = self.attraction(planet)
            total_fx += fx
            total_fy += fy
        self.xvel += total_fx / self.mass * TIMESTEP
        self.yvel += total_fy / self.mass * TIMESTEP
        self.x += self.xvel * TIMESTEP
        self.y += self.yvel * TIMESTEP
        self.distance_from_sun = math.sqrt(self.x**2 + self.y**2)
        
        self.orbit.append((self.x*SCALE+WIDTH/2,self.y*SCALE+HEIGHT/2))
        if len(self.orbit) > 100:
            self.orbit.pop(0)

        
        
    



def main():
    clock = pygame.time.Clock()
    run = True
    time_tick_ratio = FONT.render("1 tick = " + str(TIMESTEP) + " seconds", 1, (255, 255, 255))


    sun = planet("sun",0, 0, 2, 1.98892 * 10**30, (255, 255, 0),0,True)
    earth = planet("earth",x=-AU, y=0, radius=16, mass=5.9742 * 10**24, color=(0, 120, 255), yvel = 29.783 * 1000)
    mars = planet("mars",x=-AU * 1.5, y=0, radius=12, mass=6.39 * 10**23, color=(255, 0, 0), yvel = 24.077 * 1000)
    venus = planet("venus",x=-AU * 0.7, y=0, radius=14, mass=4.8685 * 10**24, color=(255, 255, 0), yvel = -35.02 * 1000)
    mercury = planet("mercury",x=-AU * 0.4, y=0, radius=8, mass=3.30 * 10**23, color=(100, 100, 100), yvel = 47.36 * 1000)
    jupiter = planet("jupiter",x=-AU * 5.2, y=0, radius=20, mass=1.898 * 10**27, color=(255, 255, 255), yvel = 13.07 * 1000)
    saturn = planet("saturn",x=-AU * 9.5, y=0, radius=16, mass=5.683 * 10**26, color=(255, 255, 0), yvel = 9.68 * 1000)

    planets = [sun,earth,mars,venus,mercury,jupiter,saturn]

    while run:
        clock.tick(60)
        screen.fill((0, 0, 0))
        screen.blit(time_tick_ratio, (10, 10))

        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                run = False
        for p in planets:
            p.update_position(planets)
            p.draw(screen)

        pygame.display.update()
    pygame.quit()

main()
