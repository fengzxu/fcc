<template>
<v-app id="inspire">
    <v-navigation-drawer v-model="drawer" app>
        <v-layout row wrap>
            <v-flex align-center justify-center layout text-xs-center>
                <img src="@/assets/guo.png">
            </v-flex>
        </v-layout>
        <v-divider />
        <v-list dense>
            <v-list-item v-for="link in menus" :key="link.text2" :to="link.route" link>
                <v-list-item-action>
                    <v-icon large color="blue darken-2">{{link.icon}}</v-icon>
                </v-list-item-action>
                <v-list-item-content>
                    <v-list-item-title>
                        <h4 class="font-weight-light">{{link.text1}}</h4>
                        <h3>{{link.text2}}</h3>
                    </v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list>
    </v-navigation-drawer>

    <v-app-bar app color="indigo" dark>
        <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
        <v-toolbar-title>基于联盟区块链的政务大数据共享</v-toolbar-title>
    </v-app-bar>

    <v-content>
        <v-container fluid>
            <router-view />
        </v-container>
    </v-content>
    <v-footer color="indigo" app>
        <span class="white--text">&copy; xujf000@gmail.com 2020</span>
    </v-footer>
</v-app>
</template>

<script>
export default {
    props: {
        source: String,
    },
    data: () => ({
        drawer: null,
        menus: [{
                icon: "fas fa-file-alt",
                text1: "住房城乡建设局",
                text2: "房屋交易备案管理",
                route: "/"
            },
            {
                icon: "fas fa-dollar-sign",
                text1: "税务局",
                text2: "不动产业务缴税",
                route: "/estatetax"
            },
            {
                icon: "fas fa-id-badge",
                text1: "自然资源和规划局",
                text2: "不动产登记办证",
                route: "/estatebook"
            },
            {
                icon: "fas fa-database",
                text1: "政务服务和大数据局",
                text2: "链上政务大数据",
                route: "/bigdata"
            },
        ],
    }),
}
</script>