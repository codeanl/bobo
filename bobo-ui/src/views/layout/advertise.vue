<template>
    <div class="box">
        <!--  -->
        <div class="search">
            <n-input round placeholder="搜索框" @keyup="handleKeyUp" />
        </div>
        <!--  -->
        <div class="banner">
            <n-carousel autoplay>
                <img class="carousel-img"
                    src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel1.jpeg" />
                <img class="carousel-img"
                    src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel2.jpeg" />
                <img class="carousel-img"
                    src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel3.jpeg" />
                <img class="carousel-img"
                    src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel4.jpeg" />
            </n-carousel>
        </div>
    </div>
</template>

<script setup lang="ts">
const handleKeyUp = () => {
    console.log("111");
};
</script>

<style scoped lang="scss">
.box {
    position: relative;

    .search {
        padding: 10px 0;
    }

    .banner {
        height: 160px;

        .carousel-img {
            height: 100%;
            width: 100%;
            object-fit: cover;
        }
    }

    .tag {
        margin-top: 10px;
        background-color: #f8f8f8;

        .title {
            p {
                color: #454545;
                font-size: 18px;
                padding: 10px 0 0 15px;
            }
        }

        .item {
            margin: 5px 15px;
            display: flex;
            justify-content: space-between;

            p:first-of-type {
                font-size: 15px;
                color: rgb(1, 116, 1);
            }

            p:nth-of-type(2) {
                color: #939191;
                font-size: 15px;
            }
        }
    }
}
</style>