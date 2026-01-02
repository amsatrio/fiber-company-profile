<template>
  <div class="min-h-screen bg-slate-50 font-sans text-slate-900">
    <nav class="sticky top-0 z-50 flex items-center justify-between bg-white/80 px-8 py-4 backdrop-blur-md shadow-sm">
      <div class="text-2xl font-bold tracking-tight text-indigo-600">TechFlow</div>
      <ul class="hidden space-x-8 font-medium md:flex">
        <li><a href="#about" class="hover:text-indigo-600 transition">About</a></li>
        <li><a href="#services" class="hover:text-indigo-600 transition">Services</a></li>
        <li><a href="#contact" class="hover:text-indigo-600 transition">Contact</a></li>
      </ul>
    </nav>

    <header class="relative overflow-hidden bg-white py-24 lg:py-32">
      <div class="container mx-auto px-6 text-center lg:text-left flex flex-col lg:flex-row items-center">
        <div class="lg:w-1/2">
          <h1 class="text-5xl font-extrabold leading-tight text-slate-900 md:text-6xl">
            We build the <span class="text-indigo-600">digital backbone</span> of your business.
          </h1>
          <p class="mt-6 text-lg text-slate-600">
            From custom SaaS platforms to high-converting landing pages, we deliver 
            engineering excellence with a human touch.
          </p>
          <div class="mt-10 flex flex-col space-y-4 sm:flex-row sm:space-x-4 sm:space-y-0 justify-center lg:justify-start">
            <button @click="scrollTo('services')" class="rounded-lg bg-indigo-600 px-8 py-4 font-bold text-white shadow-lg shadow-indigo-200 hover:bg-indigo-700">
              View Our Work
            </button>
            <button @click="scrollTo('contact')" class="rounded-lg border-2 border-slate-200 px-8 py-4 font-bold hover:bg-slate-50">
              Contact Us
            </button>
          </div>
        </div>
        <div class="mt-12 lg:mt-0 lg:w-1/2 flex justify-center">
          <div class="h-64 w-64 md:h-96 md:w-96 rounded-2xl bg-indigo-100 flex items-center justify-center text-indigo-500">
             <span class="text-6xl font-bold">TF</span>
          </div>
        </div>
      </div>
    </header>

    <section id="services" class="bg-slate-50 py-24">
      <div class="container mx-auto px-6">
        <div class="mb-16 text-center">
          <h2 class="text-3xl font-bold text-slate-900 md:text-4xl">Expertise we bring</h2>
          <p class="mt-4 text-slate-600">Tailored solutions for modern problems.</p>
        </div>
        
        <div class="grid gap-8 md:grid-cols-3">
          <div v-for="service in services" :key="service.id" 
               class="group rounded-2xl bg-white p-8 shadow-sm transition-all hover:-translate-y-2 hover:shadow-xl">
            <div class="mb-6 inline-block rounded-xl bg-indigo-50 p-4 text-3xl text-indigo-600 group-hover:bg-indigo-600 group-hover:text-white transition-colors">
              {{ service.icon }}
            </div>
            <h3 class="mb-3 text-xl font-bold">{{ service.title }}</h3>
            <p class="text-slate-600">{{ service.description }}</p>
          </div>
        </div>
      </div>
    </section>

    <section id="contact" class="py-24 bg-white">
      <div class="container mx-auto px-6 max-w-4xl">
        <div class="rounded-3xl bg-slate-900 p-8 md:p-16 text-white shadow-2xl">
          <div class="grid gap-12 md:grid-cols-2">
            <div>
              <h2 class="text-3xl font-bold">Let's talk.</h2>
              <p class="mt-4 text-slate-400">Ready to start your next project? Drop us a message and we'll get back to you within 24 hours.</p>
              <div class="mt-8 space-y-4">
                <p class="flex items-center gap-3"><span class="text-indigo-400">📍</span> San Francisco, CA</p>
                <p class="flex items-center gap-3"><span class="text-indigo-400">✉️</span> hello@techflow.io</p>
              </div>
            </div>
            
            <form @submit.prevent="handleSubmit" class="space-y-4">
              <input v-model="form.name" type="text" placeholder="Name" class="w-full rounded-lg bg-slate-800 border-none p-4 text-white focus:ring-2 focus:ring-indigo-500" />
              <input v-model="form.email" type="email" placeholder="Email" class="w-full rounded-lg bg-slate-800 border-none p-4 text-white focus:ring-2 focus:ring-indigo-500" />
              <textarea v-model="form.message" rows="4" placeholder="How can we help?" class="w-full rounded-lg bg-slate-800 border-none p-4 text-white focus:ring-2 focus:ring-indigo-500"></textarea>
              <button class="w-full rounded-lg bg-indigo-500 py-4 font-bold hover:bg-indigo-400 transition shadow-lg">Send Inquiry</button>
            </form>
          </div>
        </div>
      </div>
    </section>

    <footer class="border-t border-slate-200 py-12 text-center text-slate-500">
      <p>&copy; 2026 TechFlow Solutions. Designed with Vue & Tailwind.</p>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const services = ref([
  { id: 1, title: 'Fullstack Dev', description: 'End-to-end web applications built for scale.', icon: '⚡' },
  { id: 2, title: 'Strategy', description: 'Technical roadmapping and product consulting.', icon: '🎯' },
  { id: 3, title: 'Mobile First', description: 'Responsive designs that look great on any screen.', icon: '📱' }
]);

const form = ref({ name: '', email: '', message: '' });

const handleSubmit = () => {
  console.log('Form Submitted:', form.value);
  alert('Inquiry sent successfully!');
  form.value = { name: '', email: '', message: '' };
};
localStorage
const scrollTo = (id) => {
  document.getElementById(id).scrollIntoView({ behavior: 'smooth' });
};
</script>