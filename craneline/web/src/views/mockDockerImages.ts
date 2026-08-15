interface DockerImage {
  name: string
  tag: string
  ports: number[]
  category: 'network' | 'database' | 'app'
  description: string
}

export const mockDockerImages: DockerImage[] = [
  {
    name: 'nginx',
    tag: 'latest',
    ports: [80, 443],
    category: 'network',
    description: 'Reverse proxy / web server',
  },
  {
    name: 'postgres',
    tag: '16-alpine',
    ports: [5432],
    category: 'database',
    description: 'Base de données relationnelle',
  },
  {
    name: 'redis',
    tag: '7',
    ports: [6379],
    category: 'database',
    description: 'Cache en mémoire',
  },
  {
    name: 'traefik',
    tag: 'v3',
    ports: [80, 443, 8080],
    category: 'network',
    description: 'Reverse proxy avec dashboard',
  },
  {
    name: 'my-custom-api',
    tag: '1.4.2',
    ports: [3000],
    category: 'app',
    description: 'API interne du projet',
  },
  {
    name: 'mongo',
    tag: '7',
    ports: [27017],
    category: 'database',
    description: 'Base NoSQL orientée documents',
  },
  {
    name: 'pihole',
    tag: 'latest',
    ports: [53, 80, 443, 67],
    category: 'network',
    description: 'DNS + bloqueur de pub réseau (beaucoup de ports, pour tester une card haute)',
  },
  {
    name: 'wordpress',
    tag: 'php8.3-apache',
    ports: [80],
    category: 'app',
    description: 'CMS',
  },
]