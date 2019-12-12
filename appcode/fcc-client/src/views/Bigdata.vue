<template>
<div id="bigdata">
    <v-container grid-list-xl fluid>
        <v-tabs>
            <v-tab>
                <v-icon left>fas fa-file-alt</v-icon>
                【链上数据】房屋交易合同备案
            </v-tab>
            <v-tab>
                <v-icon left>fas fa-dollar-sign</v-icon>
                【链上数据】不动产登记业务纳税凭证
            </v-tab>
            <v-tab>
                <v-icon left>fas fa-id-badge</v-icon>
                【链上数据】不动产权证书
            </v-tab>
            <v-tab-item>
                <v-data-table :headers="hdnetcon" :items="netcons" class="elevation-1" :search="netconsearch" />
            </v-tab-item>
            <v-tab-item>
                <v-data-table :headers="hdetax" :items="etaxes" class="elevation-1" :search="etaxsearch" />
            </v-tab-item>
            <v-tab-item>
                <v-data-table :headers="hdebook" :items="ebooks" class="elevation-1" :search="ebooksearch" />
            </v-tab-item>
        </v-tabs>
    </v-container>
</div>
</template>

<script>
export default {
    data() {
        return {
            showdialog: false,
            netconsearch: '',
            etaxsearch: '',
            ebooksearch: '',
            hdnetcon: [{
                    text: '合同编号',
                    align: 'left',
                    sortable: false,
                    value: 'netconid',
                },
                {
                    text: '甲方',
                    align: 'center',
                    value: 'applya'
                },
                {
                    text: '乙方',
                    align: 'center',
                    value: 'applyb'
                },
                {
                    text: '房屋地址',
                    align: 'center',
                    value: 'addr'
                },
                {
                    text: '房屋面积',
                    align: 'center',
                    value: 'area'
                },
                {
                    text: '合同金额',
                    align: 'center',
                    value: 'balance'
                },
            ],
            hdetax: [{
                    text: '纳税凭证编号',
                    align: 'left',
                    sortable: false,
                    value: 'taxid',
                },
                {
                    text: '纳税人',
                    align: 'center',
                    value: 'taxer'
                },
                {
                    text: '房屋面积',
                    align: 'center',
                    value: 'area'
                },
                {
                    text: '纳税金额',
                    align: 'center',
                    value: 'tax'
                },
            ],
            hdebook: [{
                    text: '不动产权证书编号',
                    align: 'left',
                    sortable: false,
                    value: 'bookid',
                },
                {
                    text: '不动产所有人',
                    align: 'center',
                    value: 'owner'
                },
                {
                    text: '房屋地址',
                    align: 'center',
                    value: 'addr'
                },
                {
                    text: '房屋面积',
                    align: 'center',
                    value: 'area'
                },
            ],
            netcons: [],
            etaxes: [],
            ebooks: [],
        };
    },
    mounted() {
        this.refreshNetconList()
        this.refreshEtaxList()
        this.refreshEbookList()
    },
    methods: {
        refreshNetconList() {
            this.$axios
                .get("/api/cc/netcon/queryall", {
                    params: {},
                })
                .then(res => {
                    if (res.data.Code == 0) {
                        if (res.data.Status == null) {
                            this.netcons = []
                        } else {
                            this.netcons = res.data.Status
                        }
                    }
                })

        },
        refreshEtaxList() {
            this.$axios
                .get("/api/cc/estatetax/queryall", {
                    params: {},
                })
                .then(res => {
                    if (res.data.Code == 0) {
                        if (res.data.Status == null) {
                            this.etaxes = []
                        } else {
                            this.etaxes = res.data.Status
                        }
                    }
                })
        },
        refreshEbookList() {
            this.$axios
                .get("/api/cc/estatebook/queryall", {
                    params: {},
                })
                .then(res => {
                    if (res.data.Code == 0) {
                        if (res.data.Status == null) {
                            this.ebooks = []
                        } else {
                            this.ebooks = res.data.Status
                        }
                    }
                })

        },
    }
};
</script>
