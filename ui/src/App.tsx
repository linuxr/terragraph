import React, { useEffect, useState } from "react";
import "./App.css";
import ComboGraph from "./ComboGraph";
import { Breadcrumb, Layout, Menu } from "antd";

const { Header, Content, Footer } = Layout;

const App: React.FC = () => {
    const [data, setData] = useState({ nodes: [], edges: [], combos: [] });

    useEffect(() => {
        parseData();
    }, []);

    const parseData = async () => {
        const response = await fetch("./data.json");
        const responseData = await response.json();

        const nodes = responseData.nodes.map(
            (item: { id: string; name: string; group: string }) => ({
                id: item.id,
                label: item.name,
                comboId: item.group,
            })
        );

        const edges = responseData.edges.map(
            (item: { sourceId: string; targetId: string }) => ({
                source: item.sourceId,
                target: item.targetId,
            })
        );

        const combos = responseData.groups.map((item: string) => ({
            id: item,
            label: item,
        }));

        setData({ nodes, edges, combos });
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
                {data.nodes.length > 0 ? (
                    <div
                        className="site-layout-background"
                        style={{ padding: 24, minHeight: 380 }}
                    >
                        <ComboGraph data={data} />
                    </div>
                ) : null}
            </Content>
            <Footer style={{ textAlign: "center" }}></Footer>
        </Layout>
    );
};

export default App;
