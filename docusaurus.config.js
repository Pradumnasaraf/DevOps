// @ts-check
// `@type` JSDoc annotations allow editor autocompletion and type checking
// (when paired with `@ts-check`).
// There are various equivalent ways to declare your Docusaurus config.
// See: https://docusaurus.io/docs/api/docusaurus-config

import {themes as prismThemes} from 'prism-react-renderer';

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'DevOps',
  tagline: 'Learn DevOps',
  favicon: 'img/favicon.ico',
  url: 'https://devops.pradumnasaraf.dev',
  baseUrl: '/',
  organizationName: 'Pradumnasaraf', 
  projectName: 'DevOps',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: './sidebars.js',
          routeBasePath: '/',
          editUrl:
            'https://github.com/Pradumnasaraf/DevOps/edit/main/',
        },
        googleTagManager: {
          containerId: 'GTM-KWKGT69C',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      algolia: {
        apiKey: 'bc32ca764f4327f836b2aabf20c05424',
        indexName: 'devops',
        appId: 'ZYPCM1MC1B', // Ensure you have the correct App ID if using a multi-application setup
        contextualSearch: true, // Optional: contextual search
        // Optional: see doc section below
        // searchParameters: {},
        //... other algolia configuration
      },
      image: 'img/devops-repo-card.png',
      navbar: {
        title: 'DevOps',
        logo: {
          alt: 'DevOps Logo',
          src: 'img/logo.svg',
        },
        items: [
          {
            href: 'https://github.com/Pradumnasaraf/DevOps',
            label: 'GitHub',
            position: 'right',
          }, 
          {
            href: 'https://twitter.com/pradumna_Saraf',
            label: 'Twitter',
            position: 'right',
          },
        ],
      },
      footer: {
        logo: {
          alt: 'DevOps',
          src: 'img/logo.svg',
          href: 'https://devops.pradumnasaraf.dev',
        },
        style: 'light',
        copyright: `Copyright © ${new Date().getFullYear()} Pradumna Saraf`,
      },
      prism: {
        theme: prismThemes.github,
        darkTheme: prismThemes.dracula,
      },
    }),

  plugins: [
    [
      '@docusaurus/plugin-client-redirects',
      {
        redirects: [
          /* Argocd */
          {
            from: ['/argocd/introduction'],
            to: '/argocd',
          },
          {
            from: ['/argocd/learning-resources'],
            to: '/argocd/resources',
          },
          /* Bash Scripting */
          {
            from: ['/bash-scripting/introduction', '/bash-scripting', '/bash-introduction'],
            to: '/bash',
          },
          {
            from: ['/bash-scripting/learning-resources', '/bash-scripting/resources'],
            to: '/bash/resources',
          },
          {
            from: ['/bash-scripting/tools'],
            to: '/bash/tools',
          },
          /* DevSecOps */
          {
            from: ['/devsecops/introduction'],
            to: '/devsecops',
          },
          /* Docker */
          {
            from: ['/docker/introduction'],
            to: '/docker',
          },
          {
            from: ['/docker/learning-resources'],
            to: '/docker/resources',
          },
          /* Git */
          {
            from: ['/git/introduction'],
            to: '/git',
          },
          {
            from: ['/git/learning-resources'],
            to: '/git/resources',
          },
          /* GitHub Actions */
          {
            from: ['/github-actions/introduction'],
            to: '/github-actions',
          },
          {
            from: ['/github-actions/learning-resources'],
            to: '/github-actions/resources',
          },
          /* GitOps */
          {
            from: ['/gitops/introduction'],
            to: '/gitops',
          },
          {
            from: ['/gitops/learning-resources'],
            to: '/gitops/resources',
          },
          /* Golang */
          {
            from: ['/golang/introduction', '/go/introduction', '/go'],
            to: '/golang',
          },
          {
            from: ['/golang/learning-resources', '/go/learning-resources'],
            to: '/golang/resources',
          },
          /* Helm */
          {
            from: ['/helm/introduction'],
            to: '/helm',
          },
          {
            from: ['/helm/learning-resources'],
            to: '/helm/resources',
          },
          /* Jenkins */
          {
            from: ['/jenkins/introduction'],
            to: '/jenkins',
          },
          {
            from: ['/jenkins/learning-resources'],
            to: '/jenkins/resources',
          },
          /* Kubernetes */
          {
            from: ['/kubernetes/introduction', '/k8s/introduction', '/k8s'],
            to: '/kubernetes',
          },
          {
            from: ['/kubernetes/learning-resources', '/k8s/learning-resources'],
            to: '/kubernetes/resources',
          },
          /* Linux */
          {
            from: ['/linux/introduction'],
            to: '/linux',
          },
          {
            from: ['/linux/learning-resources'],
            to: '/linux/resources',
          },
          /* Networking */
          {
            from: ['/networking/introduction', '/network'],
            to: '/networking',
          },
          {
            from: ['/networking/learning-resources'],
            to: '/networking/resources',
          },
          /* Prometheus */
          {
            from: ['/prometheus/introduction', '/prom/introduction', '/prom'],
            to: '/prometheus',
          },
          {
            from: ['/prometheus/learning-resources'],
            to: '/prometheus/resources',
          },
          /* WebAssembly */
          {
            from: ['/webassembly/introduction', '/wasm/introduction', '/wasm'],
            to: '/webassembly',
          },
          {
            from: ['/webassembly/learning-resources', '/wasm/learning-resources'],
            to: '/webassembly/resources',
          },
          /* YAML */
          {
            from: ['/yaml/introduction'],
            to: '/yaml',
          },
          {
            from: ['/yaml/learning-resources'],
            to: '/yaml/resources',
          },
        ],
      },
    ],
  ],
};

export default config;
