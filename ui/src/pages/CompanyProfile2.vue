<template>
    <div class="min-h-screen bg-base-100 font-sans">
      
      <div class="navbar bg-base-100 shadow-lg sticky top-0 z-50 px-4 md:px-20">
        <div class="flex-1">
          <a class="text-2xl font-bold text-primary tracking-tighter cursor-pointer">NEXUS<span class="text-secondary">TECH</span></a>
        </div>
        <div class="flex-none">
          <ul class="menu menu-horizontal px-1 hidden md:flex">
            <li><a>Services</a></li>
            <li><a>Portfolio</a></li>
            <li><a>About</a></li>
          </ul>
          <button class="btn btn-primary ml-4">Contact Us</button>
        </div>
      </div>
   
      <div class="hero min-h-[70vh] bg-base-200">
        <div class="hero-content text-center">
          <div class="max-w-2xl">
            <h1 class="text-5xl font-bold">Building Digital Future</h1>
            <p class="py-6 text-lg">We create high-performance web applications and brand identities that help your business scale in the digital age.</p>
            <div class="flex gap-4 justify-center">
              <button class="btn btn-primary">Get Started</button>
              <button class="btn btn-outline">View Work</button>
            </div>
          </div>
        </div>
      </div>
   
      <div class="py-20 px-10">
        <h2 class="text-3xl font-bold text-center mb-12">Our Services</h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div class="card bg-base-100 shadow-xl border border-base-300">
            <div class="card-body items-center text-center">
              <div class="p-4 bg-primary/10 rounded-full mb-4">🚀</div>
              <h2 class="card-title">Web Development</h2>
              <p>Modern, responsive websites built with Vue and Tailwind.</p>
            </div>
          </div>
          <div class="card bg-base-100 shadow-xl border border-base-300">
            <div class="card-body items-center text-center">
              <div class="p-4 bg-secondary/10 rounded-full mb-4">🎨</div>
              <h2 class="card-title">UI/UX Design</h2>
              <p>User-centric designs that convert visitors into customers.</p>
            </div>
          </div>
          <div class="card bg-base-100 shadow-xl border border-base-300">
            <div class="card-body items-center text-center">
              <div class="p-4 bg-accent/10 rounded-full mb-4">📈</div>
              <h2 class="card-title">Cloud Solutions</h2>
              <p>Scalable infrastructure for growing enterprise needs.</p>
            </div>
          </div>
        </div>
      </div>
   
      <div class="bg-primary text-primary-content py-10">
        <div class="flex justify-center">
          <div class="stats stats-vertical lg:stats-horizontal shadow bg-transparent border-none">
            <div class="stat place-items-center">
              <div class="stat-title text-primary-content/70">Projects Done</div>
              <div class="stat-value">250+</div>
            </div>
            <div class="stat place-items-center">
              <div class="stat-title text-primary-content/70">Happy Clients</div>
              <div class="stat-value">120</div>
            </div>
            <div class="stat place-items-center">
              <div class="stat-title text-primary-content/70">Success Rate</div>
              <div class="stat-value">99%</div>
            </div>
          </div>
        </div>
      </div>
   
      <footer class="footer p-10 bg-neutral text-neutral-content mt-auto">
        <nav>
          <h6 class="footer-title">Services</h6>
          <a class="link link-hover">Branding</a>
          <a class="link link-hover">Design</a>
        </nav>
        <nav>
          <h6 class="footer-title">Company</h6>
          <a class="link link-hover">About us</a>
          <a class="link link-hover">Contact</a>
        </nav>
        <nav>
          <h6 class="footer-title">Legal</h6>
          <a class="link link-hover">Terms of use</a>
          <a class="link link-hover">Privacy policy</a>
        </nav>
      </footer>
   
    </div>
  </template>