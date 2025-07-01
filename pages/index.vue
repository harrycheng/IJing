<template>
  <div class="container">
    <header>
      <div class="logo">Welcome to IJing</div>
      <div class="description">
        IJing is a simple & powerful tool which is used to explain how everything works
      </div>
    </header>
    <div class="card">
      <div class="card-header">
        {{ result.divinatory }}
      </div>
      <div class="card-body">
        <h5 class="card-title">{{ nowDate }}</h5>
        <h5 class="card-title">{{ lunarDate }}</h5>
        <div class="card-text" v-html="result.divinatoryDetail"> </div>
      </div>
      <div class="card-footer">善易者不卜</div>
    </div>


    <footer>
      <div class="author">
        Official website:
        <a :href="`http://${website}`">{{ website }}</a> /
        Contact me:
        <a class="email" :href="`mailto:${email}`">{{ email }}</a>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick } from 'vue';
import { getLunar } from 'chinese-lunar-calendar';
import axios from 'axios';

// Define interface for divination result
interface DivinationResult {
  divinatory: string;
  divinatoryDetail: string;
}

// Define interface for the smaller cookie
interface DivinationNameCookie {
  date: string;
  name: string;
}

const result = ref<DivinationResult>({
  divinatory: 'Loading...',
  divinatoryDetail: ''
});
const nowDate = ref('');
const lunarDate = ref('');
const website = ref('IJing.com');
const email = ref('hc.harrycheng@gmail.com');

const pending = ref(false);
const error = ref<Error | null>(null);

// 获取今天的日期字符串 (YYYY-MM-DD)
const getTodayDateString = (): string => {
  const today = new Date();
  return today.toISOString().split('T')[0];
};

const fetchData = async () => {
  pending.value = true;
  error.value = null;

  try {
    const response = await axios.get('/api/divinatory');
    console.log('Axios Response:', response);

    if (response.data.success) {
      result.value = response.data.data;
      nextTick(() => {
        console.log('DOM updated with new data from API:', result.value);
      });
    } else {
      console.warn('API request was not successful.');
    }
  } catch (e) {
    error.value = e as Error;
    console.error('Error fetching data:', e);
  } finally {
    pending.value = false;
  }
};

onMounted(async () => {
  const today = new Date();
  if (isNaN(today.getTime())) {
    console.error("Invalid current date");
    return;
  }

  // Get lunar date using current date
  const lunar = getLunar(today.getFullYear(), today.getMonth() + 1, today.getDate());
  lunarDate.value = `${lunar.lunarYear} ${lunar.dateStr} `;

  // No client-side cookie check, just fetch every time.
  // The server will decide whether to return cached or new data.
  fetchData();
});
</script>

<style scoped lang="scss">
.container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

header {
  text-align: center;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  position: relative;
  padding: 2rem 0;
}

.logo {
  background-image: url('/img/tj.png');
  background-repeat: no-repeat;
  -webkit-background-size: 100px 100px;
  background-size: 100px 100px;
  background-position: center center;
  text-align: center;
  font-size: 2rem;
  padding-bottom: 120px;
  font-weight: 600;
  color: #1d1d1f;
}

.description {
  text-align: center;
  font-size: 1.1rem;
  margin-bottom: 1.5rem;
  color: #6e6e73;
}

.card {
  display: inline-block;
  width: 90%;
  max-width: 600px;
  border: 1px solid #d2d2d7;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  vertical-align: top;
  background-color: #fff;
  margin: 0 auto 20px;
}

.card-header {
  padding: 1rem 1.5rem;
  border-bottom: 1px solid #d2d2d7;
  font-size: 1.5rem;
  font-weight: 600;
  color: #1d1d1f;
}

.card-body {
  padding: 1.5rem;
  text-align: left;
}

.card-title {
  margin-top: 0;
  font-size: 1.2rem;
  font-weight: 600;
  line-height: 1.2;
  color: #1d1d1f;
}

.card-text {
  white-space: pre-line;
  margin-bottom: 1rem;
  line-height: 1.6;
  color: #333;
}

.card-footer {
  padding: 1rem 1.5rem;
  background-color: #f5f5f7;
  border-top: 1px solid #d2d2d7;
  font-style: italic;
  text-align: right;
  color: #6e6e73;
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
}

.new-btn {
  padding: 0.5rem 1rem;
  background-color: #007aff;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s;
  font-size: 1rem;
  margin-top: 20px;
}

.new-btn:hover {
  background-color: #005ecb;
}

footer {
  line-height: 1.8;
  text-align: center;
  padding: 2rem 0;
  color: #6e6e73;
  font-size: 0.9rem;
}

.backdrop {
  display: none;
}

@media (min-width: 768px) {
  .logo {
    font-size: 2.5rem;
    padding-bottom: 150px;
  }

  .description {
    font-size: 1.25rem;
    margin-bottom: 2rem;
  }

  .card {
    width: 80%;
  }

  .card-header {
    font-size: 1.75rem;
  }

  .card-title {
    font-size: 1.4rem;
  }

  footer {
    padding: 3rem 0;
  }
}

@media (min-width: 1200px) {
  .card {
    max-width: 800px;
  }
}
</style>