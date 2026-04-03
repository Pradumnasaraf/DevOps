import fs from 'node:fs';
import path from 'node:path';

const rootDir = process.cwd();
const docsDir = path.join(rootDir, 'docs');
const supportedExtensions = new Set(['.md', '.mdx']);
const ignoredPrefixes = ['http://', 'https://', 'mailto:', 'tel:', 'data:', '#', '/'];

function walk(dirPath) {
  const entries = fs.readdirSync(dirPath, {withFileTypes: true});
  const files = [];

  for (const entry of entries) {
    const fullPath = path.join(dirPath, entry.name);

    if (entry.isDirectory()) {
      files.push(...walk(fullPath));
      continue;
    }

    if (supportedExtensions.has(path.extname(entry.name))) {
      files.push(fullPath);
    }
  }

  return files;
}

function extractLinks(content) {
  const markdownLinks = [...content.matchAll(/\[[^\]]+\]\(([^)]+)\)/g)].map((match) => match[1]);
  const htmlLinks = [...content.matchAll(/(?:href|src)="([^"]+)"/g)].map((match) => match[1]);
  return [...new Set([...markdownLinks, ...htmlLinks])];
}

function normalizeTarget(target) {
  return decodeURIComponent(target.split('#')[0].split('?')[0].trim().replace(/^<|>$/g, ''));
}

function targetExists(fromFile, target) {
  const resolvedPath = path.resolve(path.dirname(fromFile), target);

  if (fs.existsSync(resolvedPath)) {
    return resolvedPath;
  }

  if (path.extname(resolvedPath)) {
    return null;
  }

  const candidates = [
    `${resolvedPath}.md`,
    `${resolvedPath}.mdx`,
    path.join(resolvedPath, 'index.md'),
    path.join(resolvedPath, 'index.mdx'),
  ];

  return candidates.find((candidate) => fs.existsSync(candidate)) ?? null;
}

const brokenLinks = [];

for (const filePath of walk(docsDir)) {
  const content = fs.readFileSync(filePath, 'utf8');

  for (const rawLink of extractLinks(content)) {
    if (ignoredPrefixes.some((prefix) => rawLink.startsWith(prefix))) {
      continue;
    }

    const target = normalizeTarget(rawLink);

    if (!target) {
      continue;
    }

    const resolvedTarget = targetExists(filePath, target);

    if (!resolvedTarget) {
      brokenLinks.push({
        filePath,
        rawLink,
        targetPath: path.resolve(path.dirname(filePath), target),
      });
    }
  }
}

if (brokenLinks.length > 0) {
  console.error('Docs contain broken local links:\n');

  for (const entry of brokenLinks) {
    const relativeFile = path.relative(rootDir, entry.filePath);
    console.error(`- ${relativeFile}: ${entry.rawLink} -> ${entry.targetPath}`);
  }

  process.exit(1);
}

console.log('Docs local links look good.');
