<template>
  <div class="layout">
    <header>
      <h1>IJing Divinatory</h1>
    </header>
    
    <main>
      <slot />
    </main>
    
    <footer>
      <p>© {{ currentYear }}  IJing Divinatory. All rights reserved.</p>
    </footer>
  </div>
</template>

<script setup lang="ts">
const currentYear = new Date().getFullYear();
</script>

<style scoped lang="scss">
.layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #f5f5f7;
  color: #1d1d1f;
}

header {
  background-color: #fff;
  color: #1d1d1f;
  padding: 1rem;
  text-align: center;
  border-bottom: 1px solid #d2d2d7;
}

nav {
  margin-top: 1rem;
}

main {
  flex: 1;
  padding: 1rem;
}

footer {
  background-color: #f5f5f7;
  padding: 1rem;
  text-align: center;
  font-size: 0.9rem;
  color: #6e6e73;
}

@media (min-width: 768px) {
  header {
    padding: 1rem 2rem;
  }

  main {
    padding: 2rem;
  }

  footer {
    padding: 1rem 2rem;
  }
}
</style>
