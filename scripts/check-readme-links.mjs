import fs from 'node:fs';
import path from 'node:path';

const rootDir = process.cwd();
const readmePath = path.join(rootDir, 'README.md');
const readme = fs.readFileSync(readmePath, 'utf8');

const markdownLinks = [...readme.matchAll(/\[[^\]]+\]\(([^)]+)\)/g)].map((match) => match[1]);
const htmlLinks = [...readme.matchAll(/(?:href|src)="([^"]+)"/g)].map((match) => match[1]);
const links = [...new Set([...markdownLinks, ...htmlLinks])];

const missing = [];

for (const link of links) {
  if (
    link.startsWith('http://') ||
    link.startsWith('https://') ||
    link.startsWith('mailto:') ||
    link.startsWith('#')
  ) {
    continue;
  }

  const cleanLink = link.split('#')[0];
  const targetPath = path.resolve(path.dirname(readmePath), cleanLink);

  if (!fs.existsSync(targetPath)) {
    missing.push({link, targetPath});
  }
}

if (missing.length > 0) {
  console.error('README contains broken local links:\n');

  for (const entry of missing) {
    console.error(`- ${entry.link} -> ${entry.targetPath}`);
  }

  process.exit(1);
}

console.log('README local links look good.');
