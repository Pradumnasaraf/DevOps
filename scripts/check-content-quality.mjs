import fs from 'node:fs';
import path from 'node:path';

const rootDir = process.cwd();
const supportedExtensions = new Set(['.md', '.mdx']);
const scanRoots = [
  path.join(rootDir, 'README.md'),
  path.join(rootDir, 'docs'),
];

const checks = [
  {regex: /Here are list of/g, message: 'Use "Here are some resources" or similar natural wording.'},
  {regex: /Check the the/g, message: 'Remove repeated words.'},
  {regex: /Wokflow files/g, message: 'Fix "Workflow" spelling.'},
  {regex: /break into seprate/g, message: 'Fix awkward phrasing and spelling.'},
  {regex: /commad name/g, message: 'Fix "command" spelling.'},
  {regex: /Disk usages capcity/g, message: 'Fix "usage/capacity" wording.'},
  {regex: /file-fath/g, message: 'Fix "file-path" spelling.'},
  {regex: /printevnv/g, message: 'Fix the command name or explanation.'},
  {regex: /the the state files/g, message: 'Remove repeated words.'},
  {regex: /\bkube exec\b/g, message: 'Use "kubectl exec" in examples.'},
  {regex: /\bkube describe\b/g, message: 'Use "kubectl describe" in examples.'},
  {regex: /\bmyngix\b/g, message: 'Fix the sample resource name typo.'},
  {regex: /\bvesrion\b/g, message: 'Fix "version" spelling.'},
  {regex: /\benvirnoment\b/g, message: 'Fix "environment" spelling.'},
  {regex: /\bavilable\b/g, message: 'Fix "available" spelling.'},
  {regex: /\bcertian\b/g, message: 'Fix "certain" spelling.'},
  {regex: /\bcareted\b/g, message: 'Fix "created" spelling.'},
  {regex: /\bchnages\b/g, message: 'Fix "changes" spelling.'},
  {regex: /\bdry nun\b/g, message: 'Fix "dry run" wording.'},
  {regex: /other Resources/g, message: 'Use consistent title casing.'},
  {regex: /^\s*docker-compose\s+/gm, message: 'Use the modern "docker compose" CLI in commands.'},
  {regex: /^\s*::set-output/gm, message: 'Use GITHUB_OUTPUT instead of ::set-output.'},
];

function walk(targetPath) {
  const stats = fs.statSync(targetPath);

  if (stats.isFile()) {
    return supportedExtensions.has(path.extname(targetPath)) ? [targetPath] : [];
  }

  const entries = fs.readdirSync(targetPath, {withFileTypes: true});
  const files = [];

  for (const entry of entries) {
    const fullPath = path.join(targetPath, entry.name);

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

const findings = [];

for (const scanRoot of scanRoots) {
  for (const filePath of walk(scanRoot)) {
    const content = fs.readFileSync(filePath, 'utf8');

    for (const check of checks) {
      for (const match of content.matchAll(check.regex)) {
        const line = content.slice(0, match.index).split('\n').length;
        findings.push({
          filePath,
          line,
          snippet: match[0],
          message: check.message,
        });
      }
    }
  }
}

if (findings.length > 0) {
  console.error('Content quality checks found issues:\n');

  for (const finding of findings) {
    const relativeFile = path.relative(rootDir, finding.filePath);
    console.error(`- ${relativeFile}:${finding.line} -> "${finding.snippet}" (${finding.message})`);
  }

  process.exit(1);
}

console.log('Content quality checks passed.');
