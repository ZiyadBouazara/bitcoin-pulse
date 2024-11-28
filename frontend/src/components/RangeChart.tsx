import React, { useEffect, useRef, useState } from 'react';
import {ColorType, createChart, IChartApi, ISeriesApi, LineData} from 'lightweight-charts';

// Define the data structure for chart data
export interface ChartData {
    time: string;
    value: number;
}

// Define the Props interface for the component
interface ChartComponentProps {
    data: Map<string, ChartData[]>;
    intervalColors: { [key: string]: string };
    initialInterval: string;
}

const ChartComponent: React.FC<ChartComponentProps> = ({ data, intervalColors, initialInterval }) => {
    const chartContainerRef = useRef<HTMLDivElement | null>(null);
    const chartRef = useRef<IChartApi | null>(null);
    const seriesRef = useRef<ISeriesApi<'Line'> | null>(null);

    // State to manage the current interval
    const [currentInterval, setCurrentInterval] = useState<string>(initialInterval);

    useEffect(() => {
        if (!chartContainerRef.current) return;

        // Create the chart instance
        const chart = createChart(chartContainerRef.current, {
            layout: {
                background: { type: ColorType.Solid, color: 'white' },
                textColor: 'black',
            },
            height: 300,
            width: chartContainerRef.current.clientWidth,
        });

        chartRef.current = chart;

        // Add a line series
        const lineSeries = chart.addLineSeries({
            color: intervalColors[initialInterval],
        });
        seriesRef.current = lineSeries;

        // Set the initial data
        lineSeries.setData(data.get(initialInterval) || []);

        // Handle resize
        const handleResize = () => {
            if (chartContainerRef.current) {
                chart.applyOptions({ width: chartContainerRef.current.clientWidth });
            }
        };

        window.addEventListener('resize', handleResize);

        return () => {
            window.removeEventListener('resize', handleResize);
            chart.remove();
        };
    }, [data, initialInterval, intervalColors]);

    // Effect to update the chart when the interval changes
    useEffect(() => {
        if (!seriesRef.current || !chartRef.current) return;

        const lineSeries = seriesRef.current;
        const chartData = data.get(currentInterval);

        if (chartData) {
            lineSeries.setData(chartData);
            lineSeries.applyOptions({ color: intervalColors[currentInterval] });
            chartRef.current.timeScale().fitContent();
        }
    }, [currentInterval, data, intervalColors]);

    return (
        <div>
            <div className="buttons-container" style={{ display: 'flex', gap: '8px' }}>
                {Array.from(data.keys()).map((interval) => (
                    <button
                        key={interval}
                        onClick={() => setCurrentInterval(interval)}
                        style={{
                            padding: '8px 16px',
                            borderRadius: '8px',
                            border: 'none',
                            cursor: 'pointer',
                            backgroundColor: currentInterval === interval ? '#d3d3d3' : '#f0f3fa',
                            color: '#131722',
                        }}
                    >
                        {interval}
                    </button>
                ))}
            </div>
            <div ref={chartContainerRef} style={{ height: '300px', marginBottom: '16px' }} />
        </div>
    );
};

export default ChartComponent;
