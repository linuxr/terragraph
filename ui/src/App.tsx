import React, { useEffect, useState } from "react";
import "./App.css";
import ComboGraph from "./ComboGraph";
import { Breadcrumb, Layout, Menu } from "antd";

const { Header, Content, Footer } = Layout;

const App: React.FC = () => {
    const [data, setData] = useState(null);

    useEffect(() => {
        getData();
    }, []);

    const getData = async () => {
        const response = await fetch("./data.json");
        const responseData = await response.json();
        setData(responseData);
    };

    return (
        <Layout>
            <Header style={{ position: "fixed", zIndex: 1, width: "100%" }}>
                <div className="logo" />
                <Menu
                    theme="dark"
                    mode="horizontal"
                    defaultSelectedKeys={["1"]}
                    items={[{ key: "1", label: "架构图" }]}
                />
            </Header>
            <Content
                className="site-layout"
                style={{ padding: "0 50px", marginTop: 64 }}
            >
                <Breadcrumb style={{ margin: "16px 0" }}></Breadcrumb>
                <div
                    className="site-layout-background"
                    style={{ padding: 24, minHeight: 380 }}
                >
                    {data ? <ComboGraph data={data} /> : ""}
                </div>
            </Content>
            <Footer style={{ textAlign: "center" }}></Footer>
        </Layout>
    );
};

export default App;
