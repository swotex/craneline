import { apiClient } from './apiClient'

export const imagesService = {
  getAll: () => apiClient.get('/images/all'),
  getById: (id) => apiClient.get(`/images/${id}`),
}