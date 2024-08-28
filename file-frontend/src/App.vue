<template>
  <v-app>
    <v-app-bar app color="blue" dark>
      <v-row justify="center" align="center" class="fill-height">
        <v-col class="text-center">
          <h1>Student Attendance Notifier</h1>
        </v-col>
      </v-row>
    </v-app-bar>

    <v-main>
      <FileUpload />
    </v-main>
  </v-app>
</template>


<script>
import FileUpload from './components/FileUpload.vue';

export default {
  name: 'App',

  components: {
    FileUpload,
  },

  data: () => ({
    //
  }),
};
</script>
