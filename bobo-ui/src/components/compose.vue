<template>
    <!-- 发布动态 -->
    <div v-if="userStore.token" style="padding: 20px;">
        <div style="display: flex;">
            <n-avatar round :size="30" :src="userStore.avatar" />
            <n-mention type="textarea" :options="options" autosize :prefix="['@', '#']" @search="handleSearch"
                :bordered="false" v-model:value="addDailyForm.content" style="margin-left: 10px;"
                placeholder="说说您的新鲜事..." @update:value="changeContent" />
        </div>
        <div style="padding: 10px 0 10px 30px">
            <n-upload ref="uploadRef" abstract list-type="image" :multiple="true" :max="9"
                action="http://localhost:4321/api/upload" :data="{
                    type: uploadType,
                }" :file-list="fileQueue" @before-upload="beforeUpload" @finish="finishUpload" @error="failUpload"
                @remove="removeUpload" @update:file-list="updateUpload">
                <div style="display: flex;align-items: center;justify-content: space-between;">
                    <!-- ICON按钮 -->
                    <div>
                        <n-upload-trigger #="{ handleClick }" abstract>
                            <n-button :disabled="videoContents.length > 0 ||
                                fileQueue.length === 9
                                " @click="() => {
                                    setUploadType('public/image');
                                    handleClick();
                                }
                                    " quaternary circle type="primary">
                                <template #icon>
                                    <n-icon size="20" color="var(--primary-color)">
                                        <image-outline />
                                    </n-icon>
                                </template>
                            </n-button>
                        </n-upload-trigger>

                        <n-upload-trigger #="{ handleClick }" abstract>
                            <n-button :disabled="imageContents.length > 0 || videoContents.length > 0 ||
                                fileQueue.length === 9
                                " @click="() => {
                                    setUploadType('public/video');
                                    handleClick();
                                }
                                    " quaternary circle type="primary">
                                <template #icon>
                                    <n-icon size="20" color="var(--primary-color)">
                                        <videocam-outline />
                                    </n-icon>
                                </template>
                            </n-button>
                        </n-upload-trigger>

                        <n-upload-trigger #="{ handleClick }" abstract>
                            <n-button :disabled="attachmentContents.length >= 1 ||
                                fileQueue.length === 9
                                " @click="() => {
                                    setUploadType('attachment');
                                    handleClick();
                                }
                                    " quaternary circle type="primary">
                                <template #icon>
                                    <n-icon size="20" color="var(--primary-color)">
                                        <attach-outline />
                                    </n-icon>
                                </template>
                            </n-button>
                        </n-upload-trigger>

                        <n-button quaternary circle type="primary" @click.stop="switchLink">
                            <template #icon>
                                <n-icon size="20" color="var(--primary-color)">
                                    <compass-outline />
                                </n-icon>
                            </template>
                        </n-button>

                        <!-- <n-button quaternary circle type="primary" @click.stop="switchEye">
                            <template #icon>
                                <n-icon size="20" color="var(--primary-color)">
                                    <eye-outline />
                                </n-icon>
                            </template>
                        </n-button> -->
                    </div>
                    <!-- 上传按钮 -->
                    <div style="display: flex;align-items: center;">
                        <n-tooltip trigger="hover" placement="bottom">
                            <template #trigger>
                                <n-progress class="text-statistic" type="circle" :show-indicator="false"
                                    status="success" :stroke-width="10"
                                    :percentage="(addDailyForm.content.length / 200) * 100"
                                    style="width: 20px; height: 20px;line-height: 20px;transform: rotate(180deg);margin-right: 5px" />
                            </template>
                            已输入{{ addDailyForm.content.length }}字
                        </n-tooltip>
                        <n-button type="primary" secondary round @click="submit" size='small'>
                            发布
                        </n-button>
                    </div>
                </div>
                <!-- 文件列表 -->
                <div>
                    <n-upload-file-list />
                </div>
            </n-upload>

            <div style="margin-top: 5px" v-if="showLinkSet">
                <n-dynamic-input v-model:value="links" placeholder="请输入以http(s)://开头的链接" :min="0" :max="3">
                    <template #create-button-default> 创建链接 </template>
                </n-dynamic-input>
            </div>

            <div class="eye-wrap" v-if="showEyeSet" style="margin-top: 5px">
                <n-radio-group v-model:value="visitType" name="radiogroup">
                    <n-space>
                        <n-radio v-for="visit in visibilities" :key="visit.value" :value="visit.value"
                            :label="visit.label" />
                    </n-space>
                </n-radio-group>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import {
    ImageOutline,
    VideocamOutline,
    AttachOutline,
    CompassOutline,
    EyeOutline,
} from '@vicons/ionicons5';
import type { MentionOption, UploadFileInfo, UploadInst } from 'naive-ui';
import useUserStore from "@/store/user";
import { computed, ref } from "vue";
let userStore = useUserStore();

import { useMessage } from 'naive-ui'
const message = useMessage()

// 发布动态表单
const addDailyForm = ref({
    content: '',
    attachment: [],
    topic: []
})

const atOptions = [
    {
        label: '07akioni',
        value: '07akioni'
    },
    {
        label: 'star-kirby',
        value: 'star-kirby'
    },
    {
        label: '广东路',
        value: '广东路'
    },
    {
        label: '颐和园路5号',
        value: '颐和园路5号'
    }
]
const sharpOptions = [
    {
        label: '它烫不了你的舌',
        value: '它烫不了你的舌'
    },
    {
        label: '也烧不了你的口',
        value: '也烧不了你的口'
    },
    {
        label: '喝醉吧',
        value: '喝醉吧'
    },
    {
        label: '不要回头',
        value: '不要回头'
    }
]
// 自定义提及 @ # 内容
const options = ref<MentionOption[]>([])
const handleSearch = (_: string, prefix: string) => {
    if (prefix === '@') {
        options.value = atOptions
    }
    else {
        options.value = sharpOptions
    }
}

// 限制输入框最大字数为200
const changeContent = (v: string) => {
    if (v.length > 200) {
        addDailyForm.value.content = v.substring(0, 200);
        message.warning("最大字数为200")
    } else {
        addDailyForm.value.content = v;
    }
}
const links = ref([])
const uploadType = ref('public/image');
const fileQueue = ref<UploadFileInfo[]>([]);
const imageContents = ref<Item.CommentItemProps[]>([]);
const videoContents = ref<Item.CommentItemProps[]>([]);
const attachmentContents = ref<Item.AttachmentProps[]>([]);
const setUploadType = (type: string) => {
    uploadType.value = type;
};

// 上传文件之前
const beforeUpload = async (data: any) => {
    // 图片类型校验
    if (
        uploadType.value === 'public/image' &&
        !['image/png', 'image/jpg', 'image/jpeg', 'image/gif'].includes(
            data.file.file?.type
        )
    ) {
        message.warning('图片仅允许 png/jpg/gif 格式');
        return false;
    }

    if (uploadType.value === 'public/image' && data.file.file?.size > 10485760) {
        message.warning('图片大小不能超过10MB');
        return false;
    }

    // 视频类型校验
    if (
        uploadType.value === 'public/video' &&
        !['video/mp4', 'video/quicktime'].includes(data.file.file?.type)
    ) {
        message.warning('视频仅允许 mp4/mov 格式');
        return false;
    }

    if (
        uploadType.value === 'public/video' &&
        data.file.file?.size > 104857600
    ) {
        message.warning('视频大小不能超过100MB');
        return false;
    }
    // 附件类型校验
    if (
        uploadType.value === 'attachment' && !(await isZipFile(data.file.file))
    ) {
        message.warning('附件仅允许 zip 格式');
        return false;
    }

    if (uploadType.value === 'attachment' && data.file.file?.size > 104857600) {
        message.warning('附件大小不能超过100MB');
        return false;
    }

    return true;
};

const finishUpload = ({ file, event }: any): any => {
    try {
        let data = JSON.parse(event.target?.response);

        if (data.code === 200) {
            if (uploadType.value === 'public/image') {
                imageContents.value.push({
                    id: file.id,
                    content: data.data
                } as Item.CommentItemProps);
            }
            if (uploadType.value === 'public/video') {
                videoContents.value.push({
                    id: file.id,
                    content: data.data,
                } as Item.CommentItemProps);
            }
            if (uploadType.value === 'attachment') {
                attachmentContents.value.push({
                    id: file.id,
                    content: data.data,
                } as Item.AttachmentProps);
            }
        }
    } catch (error) {
        message.error('上传失败');
    }
};

const failUpload = ({ file, event }: any): any => {
    try {
        let data = JSON.parse(event.target?.response);
        if (data.code !== 200) {
            let errMsg = data.message || '上传失败';
            if (data.details && data.details.length > 0) {
                data.details.map((detail: string) => {
                    errMsg += ':' + detail;
                });
            }
            message.error(errMsg);
        }
    } catch (error) {
        message.error('上传失败');
    }
};

const removeUpload = ({ file }: any) => {
    let idx = imageContents.value.findIndex((item) => item.id === file.id);
    if (idx > -1) {
        imageContents.value.splice(idx, 1);
    }
    idx = videoContents.value.findIndex((item) => item.id === file.id);
    if (idx > -1) {
        videoContents.value.splice(idx, 1);
    }
    idx = attachmentContents.value.findIndex((item) => item.id === file.id);
    if (idx > -1) {
        attachmentContents.value.splice(idx, 1);
    }
};

const updateUpload = (list: UploadFileInfo[]) => {
    for (let i = 0; i < list.length; i++) {
        var name = list[i].name;
        var basename: string = name.split('.').slice(0, -1).join('.');
        var ext: string = name.split('.').pop()!;
        if (basename.length > 30) {
            list[i].name = basename.substring(0, 18) + "..." + basename.substring(basename.length - 9) + "." + ext;
        }
    }
    fileQueue.value = list;
    console.log(fileQueue.value);
};

//判断文件是否为压缩包
const isZipFile = (file: File): Promise<unknown> => {
    const fileReader = new FileReader();

    const isValidZipFileType = (fileType: string): boolean => {
        const zipFileTypes = ['application/zip', 'application/x-zip', 'application/octet-stream', 'application/x-zip-compressed'];
        return zipFileTypes.includes(fileType);
    };

    const checkFileType = (): boolean => {
        const arr = new Uint8Array(fileReader.result as ArrayBuffer).subarray(0, 4);
        let header = '';
        for (let i = 0; i < arr.length; i++) {
            header += arr[i].toString(16);
        }

        switch (header) {
            case '504b0304':
            case '504b0506':
            case '504b0708':
                return isValidZipFileType('application/zip');
            case '504b030414':
                return isValidZipFileType('application/x-zip-compressed');
            case '504b0508':
                return isValidZipFileType('application/x-zip');
            case '504b5370':
                return isValidZipFileType('application/octet-stream');
            default:
                return false;
        }
    };

    return new Promise((resolve, reject) => {
        fileReader.onloadend = () => {
            const fileType = file.type;
            if (fileType === '' || fileType === 'application/octet-stream') {
                // 如果浏览器不能识别文件类型，则进行手动检查
                resolve(checkFileType());
            } else {
                // 如果浏览器可以识别文件类型，则根据文件类型进行检查
                resolve(isValidZipFileType(fileType));
            }
        };
        fileReader.readAsArrayBuffer(file.slice(0, 4));
    });
}
const showLinkSet = ref(false)
const showEyeSet = ref(false)
const switchLink = () => {
    showLinkSet.value = !showLinkSet.value;
    if (showLinkSet.value && showEyeSet.value) {
        showEyeSet.value = false
    }
};

const visitType = ref<any>('PUBLIC');
const visibilities = computed(() => {
    let res = [
        { value: 'PUBLIC', label: "公开" },
        { value: 'PRIVATE', label: "私密" },
        { value: 'Following', label: "关注可见" },
    ];
    if (userStore.token) {
        res.push({ value: 'FRIEND', label: "好友可见" });
    }
    return res;
});

const switchEye = () => {
    showEyeSet.value = !showEyeSet.value;
    if (showEyeSet.value && showLinkSet.value) {
        showLinkSet.value = false
    }
};

// 
import { defineEmits } from 'vue';
const emit = defineEmits(['submit-success']);
import { SaveOrUpdate } from '@/api/daily'
const submit = async () => {
    if (addDailyForm.value.content) {

        imageContents.value.forEach(item => {
            addDailyForm.value.attachment.push({
                url: item.content,
                type: '1'
            })
        })
        videoContents.value.forEach(item => {
            addDailyForm.value.attachment.push({
                url: item.content,
                type: '2'
            })
        })
        attachmentContents.value.forEach(item => {
            addDailyForm.value.attachment.push({
                url: item.content,
                type: '3'
            })
        })
        links.value.forEach(item => {
            addDailyForm.value.attachment.push({
                url: item,
                type: '4'
            })
        })
        console.log(addDailyForm.value);

        let res = await SaveOrUpdate(addDailyForm.value)
        if (res.code == 200) {
            message.success("发表成功")
            addDailyForm.value = {
                content: '',
                attachment: [],
                topic: []
            }

            fileQueue.value = []
            links.value = []
            imageContents.value = []
            videoContents.value = []
            attachmentContents.value = []
            // 
            showLinkSet.value = false
            showEyeSet.value = false
            // 
            emit('submit-success');
        } else {
            message.error(res.message)
        }
    } else {
        message.warning("请先输入内容哦")
    }
}

</script>

<style scoped></style>