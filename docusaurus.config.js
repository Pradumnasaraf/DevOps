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
          {
            from: ['/webassembly/introduction'], // Old paths
            to: '/webassembly', // Target path
          },
        ],
      },
    ],
  ],
};

export default config;
