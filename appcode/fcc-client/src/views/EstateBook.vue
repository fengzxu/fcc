<template>
<div id="netcon">
    <v-container grid-list-xl fluid>
        <v-card class="mx-auto" max-width="100%">
            <v-divider></v-divider>
            <v-card>
                <v-card-title class="grey lighten-4">不动产登记业务受理与证书发放
                    <v-spacer></v-spacer>
                    <v-text-field v-model="search" label="按内容查询" single-line hide-details></v-text-field>
                    <v-btn icon @click="refreshBooks">
                        <v-icon>fas fa-sync-alt</v-icon>
                    </v-btn>
                    <v-btn icon @click="onCreate">
                        <v-icon>fas fa-plus</v-icon>
                    </v-btn>
                </v-card-title>
                <v-data-table :headers="headers" :items="items" class="elevation-1" :search="search">
                    <template v-slot:item.BookID="{ item }">
                        <v-chip :color="getColorByCC(item.IsCCed)">{{ item.BookID }}</v-chip>
                    </template>
                    <template v-slot:item.action="{ item }">
                        <v-btn color="success" :disabled="!getDisToCC(item.IsCCed)" @click="onToCC(item.ID,item.BookID,item.NetconID,item.TaxID,item.Owner,item.Addr,item.Area)">上链</v-btn>
                        <v-btn color="info" :disabled="getDisToCC(item.IsCCed)" @click="queryCC(item.BookID)">查链</v-btn>
                    </template>
                </v-data-table>
            </v-card>
        </v-card>
        <v-dialog v-model="showdialog" width="800">
            <v-card fill-height>
                <v-card-title class="title grey lighten-2" primary-title>{{getDialogTitle(dialogtype)}}</v-card-title>
                <v-card-text>
                    <v-row>
                        <v-col cols="12">
                            <v-text-field v-model="ebook.id" :hidden="true" />
                            <v-text-field v-model="ebook.bookid" label="不动产权证书编号" placeholder="" :rules="[rules.required]" :autofocus=true :readonly="getReadonly()" />
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col cols="6">
                            <v-text-field v-model="ebook.owner" label="产权所有人" placeholder="" :rules="[rules.required]" :readonly="getReadonly()" />
                        </v-col>
                        <v-col cols="6">
                            <v-text-field v-model="ebook.area" label="房屋面积" placeholder="" :rules="[rules.required]" :readonly="getReadonly()" />
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col cols="12">
                            <v-text-field v-model="ebook.addr" label="房屋地址" placeholder="" :rules="[rules.required]" :readonly="getReadonly()" />
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col cols="6">
                            <v-select v-model="ebook.netconid" :items="netconids" label="合同备案号(链上查询)" placeholder="" :readonly="getReadonly()" />
                        </v-col>
                        <v-col cols="6">
                            <v-select v-model="ebook.taxid" :items="etaxids" label="缴纳凭证号(链上查询)" placeholder="" :readonly="getReadonly()" />
                        </v-col>
                    </v-row>
                </v-card-text>
                <v-card-actions justify-right>
                    <v-btn color="primary" rounded @click="toCC()" :disabled="!showToCC">上链</v-btn>
                    <v-btn color="primary" rounded @click="saveNew()" :disabled="!showSave">保存</v-btn>
                    <v-btn color="warning" rounded @click="showdialog = false">关闭</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-snackbar v-model="sb.show" :color="sb.color" :timeout="sb.timeout" :top="true">{{sb.text}}</v-snackbar>
    </v-container>
</div>
</template>

<script>
export default {
    data() {
        return {
            showdialog: false,
            showSave: false,
            showToCC: false,
            dialogtype: 1,
            ebook: {},
            netconids: [],
            etaxids: [],
            search: '',
            sb: { //snakebar
                show: false,
                color: "",
                text: "",
                timeout: 3000,
            },
            headers: [{
                    text: '办理时间',
                    align: 'left',
                    sortable: false,
                    value: 'CreateDT',
                },
                {
                    text: '不动产权证书编号',
                    align: 'left',
                    sortable: false,
                    value: 'BookID',
                },
                {
                    text: '产权所有人',
                    align: 'center',
                    value: 'Owner'
                },
                {
                    text: '房屋地址',
                    align: 'center',
                    value: 'Addr'
                },
                {
                    text: '房屋面积',
                    align: 'center',
                    value: 'Area'
                },
                {
                    text: '合同备案号',
                    align: 'center',
                    value: 'NetconID'
                },
                {
                    text: '纳税凭证号',
                    align: 'center',
                    value: 'TaxID'
                },
                {
                    text: '数据上链',
                    align: 'center',
                    sortable: false,
                    value: 'action'
                },
            ],
            items: [],
            rules: {
                required: value => !!value || '请输入',
                counter: value => value.length <= 20 || '最多20个字符',
            },
        };
    },
    mounted() {
        this.refreshBooks()
        this.refreshNetconids()
        this.refreshEtaxids()
    },
    methods: {
        refreshBooks() {
            this.$axios
                .get("/api/estatebook/queryall", {
                    params: {},
                })
                .then(res => {
                    if (res.data.Code == 0) {
                        if (res.data.Status == null) {
                            this.items = []
                        } else {
                            this.items = res.data.Status
                        }
                    }
                })
        },
        refreshNetconids() {
            this.$axios
                .get("/api/cc/netcon/queryallid", {
                    params: {},
                })
                .then(res => {
                    if (res.data.Code == 0) {
                        if (res.data.Status == null) {
                            this.netconids = []
                        } else {
                            this.netconids = res.data.Status
                        }
                    }
                })
        },
        refreshEtaxids() {
            this.$axios
                .get("/api/cc/estatetax/queryallid", {
                    params: {},
                })
                .then(res => {
                    if (res.data.Code == 0) {
                        if (res.data.Status == null) {
                            this.etaxids = []
                        } else {
                            this.etaxids = res.data.Status
                        }
                    }
                })
        },
        saveNew() {
            let data = {
                bookid: this.ebook.bookid,
                netconid: this.ebook.netconid,
                taxid: this.ebook.taxid,
                owner: this.ebook.owner,
                addr: this.ebook.addr,
                area: this.ebook.area,
            }
            this.$axios
                .post("/api/estatebook/create", this.$qs.stringify(data))
                .then(res => {
                    if (res.data.Code == 0) {
                        this.refreshBooks()
                        this.showdialog = false
                        this.ebook = {}
                        this.sb.color = "success"
                        this.sb.text = "操作成功!"
                        this.sb.show = true
                    } else {
                        this.sb.color = "error"
                        this.sb.text = "操作失败：" + res.data.Status
                        // this.sb.show = true
                    }
                })
        },
        getColorByCC(cc) {
            if (cc > 0)
                return "green";
            return "warning";
        },
        getDisToCC(code) {
            if (code > 0)
                return false;
            return true;
        },
        getDialogTitle(dtype) {
            switch (dtype) {
                case 1:
                    return '新增不动产权证书'
                case 2:
                    return '不动产权证书上链'
                case 3:
                    return '不动产权证书链上查询'
            }
        },
        onCreate() {
            this.dialogtype = 1
            this.showSave = true
            this.showToCC = false
            this.ebook = {}
            this.showdialog = true
        },
        onToCC(uuid, bookid, netconid, taxid, owner, addr, area) {
            this.dialogtype = 2
            this.showSave = false
            this.showToCC = true
            this.ebook.uuid = uuid
            this.ebook.bookid = bookid
            this.ebook.netconid = netconid
            this.ebook.taxid = taxid
            this.ebook.owner = owner
            this.ebook.addr = addr
            this.ebook.area = area
            this.showdialog = true
        },
        toCC() {
            let data = {
                uuid: this.ebook.uuid,
                bookid: this.ebook.bookid,
                netconid: this.ebook.netconid,
                taxid: this.ebook.taxid,
                owner: this.ebook.owner,
                area: this.ebook.area,
                addr: this.ebook.addr,
            }
            this.$axios
                .post("/api/estatebook/tocc", this.$qs.stringify(data))
                .then(res => {
                    if (res.data.Code == 0) {
                        this.showdialog = false
                        this.ebook = {}
                        this.sb.color = "success"
                        this.sb.text = "操作成功,不动产权证书已上链！"
                        this.refreshBooks()
                        this.sb.show = true
                    } else {
                        this.sb.color = "error"
                        this.sb.text = "操作失败：" + res.data.Status
                        this.sb.show = true
                    }
                })
        },
        queryCC(bookid) {
            this.showdialog = true
            this.$axios
                .get("/api/cc/estatebook/querybybookid?bookid=" + bookid)
                .then(res => {
                    if (res.data.Code == 0) {
                        this.showSave = false
                        this.showToCC = false
                        this.dialogtype = 3
                        this.ebook = res.data.Status[0]
                    } else {
                        this.showdialog = false
                        this.sb.color = "error"
                        this.sb.text = "查询链上不动产权证书失败：" + res.data.Status
                        this.sb.show = true
                    }
                })
        },
        getReadonly() {
            switch (this.dialogtype) {
                case 1:
                    return false
                case 2:
                    return true
                case 3:
                    return true
            }
        },
    }
};
</script>
