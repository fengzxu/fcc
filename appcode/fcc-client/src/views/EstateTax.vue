<template>
<div id="netcon">
    <v-container grid-list-xl fluid>
        <v-card class="mx-auto" max-width="100%">
            <v-divider></v-divider>
            <v-card>
                <v-card-title class="grey lighten-4">不动产登记业务缴税
                    <v-spacer></v-spacer>
                    <v-text-field v-model="search" label="按内容查询" single-line hide-details></v-text-field>
                    <v-btn icon @click="refreshTaxList">
                        <v-icon>fas fa-sync-alt</v-icon>
                    </v-btn>
                    <v-btn icon @click="onCreate">
                        <v-icon>fas fa-plus</v-icon>
                    </v-btn>
                </v-card-title>
                <v-data-table :headers="headers" :items="items" class="elevation-1" :search="search">
                    <template v-slot:item.TaxID="{ item }">
                        <v-chip :color="getColorByCC(item.IsCCed)">{{ item.TaxID }}</v-chip>
                    </template>
                    <template v-slot:item.action="{ item }">
                        <v-btn color="success" :disabled="!getDisToCC(item.IsCCed)" @click="onToCC(item.ID,item.TaxID,item.Taxer,item.Area,item.Tax)">上链</v-btn>
                        <v-btn color="info" :disabled="getDisToCC(item.IsCCed)" @click="queryCC(item.TaxID)">查链</v-btn>
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
                            <v-text-field v-model="etax.id" :hidden="true" />
                            <v-text-field v-model="etax.taxid" label="缴税凭证号" placeholder="" :rules="[rules.required]" :autofocus=true :readonly="getReadonly()" />
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col cols="12">
                            <v-text-field v-model="etax.taxer" label="纳税人" placeholder="" :rules="[rules.required]" :readonly="getReadonly()" />
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col cols="6">
                            <v-text-field v-model="etax.area" label="房屋面积" placeholder="" :rules="[rules.required]" :readonly="getReadonly()" />
                        </v-col>
                        <v-col cols="6">
                            <v-text-field v-model="etax.tax" label="缴税金额" placeholder="" :rules="[rules.required]" :readonly="getReadonly()" />
                        </v-col>
                    </v-row>
                </v-card-text>
                <v-card-actions justify-right>
                    <v-btn color="primary" rounded @click="toCC()" :disabled="!showToCC">上链</v-btn>
                    <v-btn color="primary" rounded @click="saveNew()" :disabled="!showSave">保存</v-btn>
                    <v-btn color="warning" rounded @click="etax='';showdialog = false">关闭</v-btn>
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
            etax: {},
            search: '',
            sb: { //snakebar
                show: false,
                color: "",
                text: "",
                timeout: 3000,
            },
            headers: [{
                    text: '缴税时间',
                    align: 'left',
                    sortable: false,
                    value: 'CreateDT',
                },
                {
                    text: '纳税凭证编号',
                    align: 'left',
                    sortable: false,
                    value: 'TaxID',
                },
                {
                    text: '纳税人',
                    align: 'center',
                    value: 'Taxer'
                },
                {
                    text: '房屋面积',
                    align: 'center',
                    value: 'Area'
                },
                {
                    text: '纳税金额',
                    align: 'center',
                    value: 'Tax'
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
        this.refreshTaxList()
    },
    methods: {
        refreshTaxList() {
            this.$axios
                .get("/api/estatetax/queryall", {
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
        saveNew() {
            let data = {
                taxid: this.etax.taxid,
                taxer: this.etax.taxer,
                area: this.etax.area,
                tax: this.etax.tax,
            }
            this.$axios
                .post("/api/estatetax/create", this.$qs.stringify(data))
                .then(res => {
                    if (res.data.Code == 0) {
                        this.refreshTaxList()
                        this.showdialog = false
                        this.etax = {}
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
                    return '新增纳税凭证'
                case 2:
                    return '纳税凭证上链'
                case 3:
                    return '链上纳税凭证查询'
            }
        },
        onCreate() {
            this.dialogtype = 1
            this.showSave = true
            this.showToCC = false
            this.etax = {}
            this.showdialog = true
        },
        onToCC(uuid, taxid, taxer, area, tax) {
            this.dialogtype = 2
            this.showSave = false
            this.showToCC = true
            this.etax.uuid = uuid
            this.etax.taxid = taxid
            this.etax.taxer = taxer
            this.etax.area = area
            this.etax.tax = tax
            this.showdialog = true
        },
        toCC() {
            let data = {
                uuid: this.etax.uuid,
                taxid: this.etax.taxid,
                taxer: this.etax.taxer,
                area: this.etax.area,
                tax: this.etax.tax,
            }
            this.$axios
                .post("/api/estatetax/tocc", this.$qs.stringify(data))
                .then(res => {
                    if (res.data.Code == 0) {
                        this.showdialog = false
                        this.etax = {}
                        this.sb.color = "success"
                        this.sb.text = "操作成功,纳税凭证已上链！"
                        this.refreshTaxList()
                        this.sb.show = true
                    } else {
                        this.sb.color = "error"
                        this.sb.text = "操作失败：" + res.data.Status
                        this.sb.show = true
                    }
                })
        },
        queryCC(taxid) {
            this.showdialog = true
            this.$axios
                .get("/api/cc/estatetax/querybytaxid?taxid=" + taxid)
                .then(res => {
                    if (res.data.Code == 0) {
                        this.showSave = false
                        this.showToCC = false
                        this.dialogtype = 3
                        this.etax = res.data.Status[0]
                    } else {
                        this.showdialog = false
                        this.sb.color = "error"
                        this.sb.text = "查询链上纳税凭证失败：" + res.data.Status
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
