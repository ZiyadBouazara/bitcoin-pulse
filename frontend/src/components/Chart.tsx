import { createChart, ColorType, IChartApi, ISeriesApi } from 'lightweight-charts';
import React, { useEffect, useRef } from 'react';

export interface ChartData {
    time: string;
    value: number;
}

export const colors = {
    backgroundColor: 'white',
    lineColor:'#2962FF',
    textColor: 'black',
    areaTopColor: '#2962FF',
    areaBottomColor: 'rgba(41, 98, 255, 0.28)',
}

interface Colors {
    backgroundColor?: string;
    lineColor?: string;
    textColor?: string;
    areaTopColor?: string;
    areaBottomColor?: string;
}
interface Props {
    data: ChartData[];
    colors: Colors;
}

function ChartComponent({data, colors}:Props) {
    // Ensure the ref has a correct type
    const chartContainerRef = useRef<HTMLDivElement | null>(null);
    const chartRef = useRef<IChartApi | null>(null);
    const seriesRef = useRef<ISeriesApi<'Area'> | null>(null);

    useEffect(() => {
        if (!chartContainerRef.current) return;

        // Create the chart instance
        const chart = createChart(chartContainerRef.current, {
            layout: {
                background: { type: ColorType.Solid, color: colors.backgroundColor },
                textColor: colors.textColor,
            },
            width: chartContainerRef.current.clientWidth,
            height: 300,
        });

        chartRef.current = chart;

        // Add area series
        const areaSeries = chart.addAreaSeries({
            lineColor: colors.lineColor,
            topColor: colors.areaTopColor,
            bottomColor: colors.areaBottomColor,
        });

        seriesRef.current = areaSeries;

        // Set the data
        areaSeries.setData(data);

        // Handle resizing
        const handleResize = () => {
            if (chartContainerRef.current) {
                chart.applyOptions({ width: chartContainerRef.current.clientWidth });
            }
        };

        window.addEventListener('resize', handleResize);

        // Fit content initially
        chart.timeScale().fitContent();

        // Cleanup on unmount
        return () => {
            window.removeEventListener('resize', handleResize);

            // Remove chart instance properly
            chart.remove();
            chartRef.current = null;
            seriesRef.current = null;
        };
    }, [data, colors.backgroundColor, colors.lineColor, colors.textColor, colors.areaTopColor, colors.areaBottomColor]);

    return <div ref={chartContainerRef} style={{ position: 'relative', height: '300px' }} />;
}

export default ChartComponent;