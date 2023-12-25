<template>
  <div class="wrapper">
    <div class="item">
      <q-form @submit.prevent="onSubmit" >
          <q-card>
            <q-card-section>
              <div class="text-h6">Войдите в систему</div>
            </q-card-section>
            <q-separator />
            <q-card-actions vertical>
              <q-input v-model.trim="login" label="Логин" type="text" required />
              <q-input v-model.trim="password" label="Пароль" type="password" required />
            </q-card-actions>
          </q-card>
        <q-btn type="submit" label="Войти" class="q-ma-md custom__btn" />
      </q-form>
    </div>
  </div>
</template>
  
<script>
import { defineComponent, onMounted } from 'vue';

export default defineComponent({
  name: 'LoginPage',

  data() {
    return {
      login: '',
      password: '',
    };
  },

  setup() {
    onMounted(() => {
    });
  },

  methods: {
    onSubmit() {
      // TODO: Обработать ввод пользователя и выполнить вход
    },
  },
});
</script>

<style lang="scss">
.wrapper {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  grid-gap: 10px;
  grid-auto-rows: 400px;
  grid-template-areas:
    ". a a ."
    ". a a .";
}

.item {
  grid-area: a;
  align-self: center;
}

.custom__btn{
  width: 100%;
  max-width: 200px;
}
</style>