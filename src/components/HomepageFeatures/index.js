import Link from '@docusaurus/Link';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

const topics = [
  {
    title: 'Docker',
    href: '/docker',
    description: 'Containers, images, Compose, and packaging fundamentals.',
  },
  {
    title: 'Kubernetes',
    href: '/kubernetes',
    description: 'Architecture, workloads, services, and day-to-day cluster concepts.',
  },
  {
    title: 'GitHub Actions',
    href: '/github-actions',
    description: 'Workflow basics, runners, triggers, and reusable automation patterns.',
  },
  {
    title: 'Terraform',
    href: '/terraform',
    description: 'Infrastructure as code, providers, plans, and dependency management.',
  },
];

export default function HomepageFeatures() {
  return (
    <section className={styles.topicsSection}>
      <div className={styles.header}>
        <Heading as="h2" className={styles.title}>
          Explore Core Topics
        </Heading>
        <p className={styles.description}>
          Start with the areas most people reach for first, then expand into the rest of the
          library from the sidebar.
        </p>
      </div>
      <div className={styles.grid}>
        {topics.map((topic) => (
          <Link key={topic.title} className={styles.card} to={topic.href}>
            <Heading as="h3" className={styles.cardTitle}>
              {topic.title}
            </Heading>
            <p className={styles.cardDescription}>{topic.description}</p>
          </Link>
        ))}
      </div>
    </section>
  );
}
