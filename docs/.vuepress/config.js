const glob = require("glob");
const markdownIt = require("markdown-it");
const meta = require("markdown-it-meta");
const fs = require("fs");
const _ = require("lodash");

const sidebar = (directory, array) => {
    return array.map(i => {
        const children = _.sortBy(
            glob
                .sync(`./${directory}/${i[1]}/*.md`)
                .map(path => {
                    const md = new markdownIt();
                    const file = fs.readFileSync(path, "utf8");
                    md.use(meta);
                    md.render(file);
                    const order = md.meta.order;
                    return { path, order };
                })
                .filter(f => f.order !== false),
            ["order", "path"]
        )
            .map(f => f.path)
            .filter(f => !f.match("README"));

        return {
            title: i[0],
            children
        };
    });
};

module.exports = {
    base: "/plugchain/",
    plugins: [
        ['@vuepress/search', {
            searchMaxSuggestions: 10
        }]
    ],
    locales: {
        "/": {
            lang: "en-US",
            title: "Plug Chain Documents",
            description: "Plug Chain Documents",
        },
        "/zh/": {
            lang: "简体中文",
            title: "Plug Chain 文档",
            description: "Plug Chain 文档",
        }
    },
    themeConfig: {
        locales: {
            "/": {
                selectText: 'Languages',
                label: 'English',
                nav: [
                    {
                        text: 'Back to Plug Chain',
                        link: 'https://en.plugchain.info/'
                    },
                    {
                        text:'Github',
                        link: 'https://github.com/oracleNetworkProtocol/plugchain'
                    },
                    {
                        text: 'Jion us',
                        ariaLabel: 'Community',
                        items: [
                          { text: 'Telegram', link: 'https://t.me/plugchain', target:'_blank' },
                          { text: 'Twitter', link: 'https://twitter.com/Plugchainclub',target:'_blank' }
                        ]
                    }
                    
                ],
                sidebar: sidebar("", [
                    ["Getting Started", "get-started"],
                    ["Concepts", "concepts"],
                    ["Features", "features"],
                    ["Daemon", "daemon"],
                    ["CLI Client", "cli-client"],
                    ["Endpoints", "endpoints"],
                    ["Tools", "tools"],
                    ["Migration", "migration"],
                    ["Evm", "evm"]
                ])
            },
            "/zh/": {
                selectText: '选择语言',
                label: '简体中文',
                nav: [
                    {
                        text: 'Plug Chain 官网',
                        link: 'https://zh.plugchain.info/'
                    },
                    {
                        text:'Github',
                        link: 'https://github.com/oracleNetworkProtocol/plugchain'
                    },
                    {
                        text: '加入我们',
                        ariaLabel: '社区',
                        items: [
                          { text: 'Telegram', link: 'https://t.me/plugchain', target:'_blank' },
                          { text: 'Twitter', link: 'https://twitter.com/Plugchainclub',target:'_blank' }
                        ]
                    }
                ],
                sidebar: sidebar("", [
                    ["快速开始", "/zh/get-started"],
                    ["概念", "/zh/concepts"],
                    ["功能模块", "/zh/features"],
                    ["守护进程", "/zh/daemon"],
                    ["命令行客户端", "/zh/cli-client"],
                    ["服务端点", "/zh/endpoints"],
                    ["工具", "/zh/tools"],
                    ["迁移", "/zh/migration"],
                    ["evm", "/zh/evm"]
                ])
            }
        },
    }
};