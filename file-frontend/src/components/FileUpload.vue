<template>
  <v-container>
    <v-card
      class="mx-auto"
      max-width="600"
      outlined
    >
      <v-form @submit.prevent="uploadFile">
        <v-card-text>Upload Attendance File:</v-card-text>
        <v-file-input v-model="file" label="Upload Excel File" prepend-icon="mdi-upload" accept=".xlsx" show-size
          required></v-file-input>

        <v-btn color="primary" type="submit">Upload and Send Emails</v-btn>
      </v-form>

    </v-card>
  </v-container>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      file: null,
    };
  },
  methods: {
    async uploadFile() {
      if (!this.file) {
        this.$toast.error("Please select a file to upload.");
        return;
      }

      const formData = new FormData();
      formData.append("file", this.file);

      try {
        const response = await axios.post("http://localhost:8080/api/upload", formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });
        console.log(response.data);
        this.$toast.success("File uploaded and emails sent successfully!");
      } catch (error) {
        this.$toast.error("Failed to upload file. Please try again.");
        console.error("File upload error:", error);
      }
    },
  },
};
</script>

<style scoped>
.v-container {
  max-width: 600px;
  margin: auto;
  padding-top: 50px;
}
</style>
