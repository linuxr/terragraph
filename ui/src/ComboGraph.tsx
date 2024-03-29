import React, { useEffect } from "react";
import G6 from "@antv/g6";

const ComboGraph = (props: { data: any }) => {
  const { data } = props;

  useEffect(() => {
    const container = document.getElementById("container");
    const width = container?.scrollWidth;
    const height = (container?.scrollHeight || 500) - 30;

    const graph = new G6.Graph({
      container: "container",
      width,
      height: height - 50,
      fitView: true,
      fitViewPadding: 30,
      animate: true,
      groupByTypes: false,
      modes: {
        default: [
          "drag-combo",
          "zoom-canvas",
          // "drag-node",
          "drag-canvas",
          {
            type: "collapse-expand-combo",
            relayout: false,
          },
        ],
      },
      layout: {
        type: "dagre",
        sortByCombo: false,
        ranksep: 10,
        nodesep: 10,
      },
      defaultNode: {
        size: [60, 30],
        type: "rect",
        anchorPoints: [
          [0.5, 0],
          [0.5, 1],
        ],
      },
      defaultEdge: {
        type: "line",
        style: {
          radius: 20,
          offset: 45,
          endArrow: true,
          lineWidth: 1,
          stroke: "#C2C8D5",
        },
      },
      defaultCombo: {
        type: "rect",
        style: {
          fillOpacity: 0.1,
        },
      },
    });

    console.log(data);
    graph.data(data);
    graph.render();
  });

  return (
    <div
      id="container"
      className="site-layout-background"
      style={{ padding: 24, minHeight: 600 }}
    />
  );
};

export default ComboGraph;
