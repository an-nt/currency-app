USE [master]
GO
/****** Object:  Database [Supermarket]    Script Date: 10/7/2020 5:34:59 PM ******/
CREATE DATABASE [Supermarket]
 CONTAINMENT = NONE
 ON  PRIMARY 
( NAME = N'Supermarket1', FILENAME = N'D:\Code\Microsoft SQL Server\Supermarket\Supermarket1.mdf' , SIZE = 8192KB , MAXSIZE = UNLIMITED, FILEGROWTH = 65536KB )
 LOG ON 
( NAME = N'Supermarket1_log', FILENAME = N'D:\Code\Microsoft SQL Server\Supermarket\Supermarket1_log.ldf' , SIZE = 8192KB , MAXSIZE = 2048GB , FILEGROWTH = 65536KB )
GO
ALTER DATABASE [Supermarket] SET COMPATIBILITY_LEVEL = 140
GO
IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [Supermarket].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO
ALTER DATABASE [Supermarket] SET ANSI_NULL_DEFAULT OFF 
GO
ALTER DATABASE [Supermarket] SET ANSI_NULLS OFF 
GO
ALTER DATABASE [Supermarket] SET ANSI_PADDING OFF 
GO
ALTER DATABASE [Supermarket] SET ANSI_WARNINGS OFF 
GO
ALTER DATABASE [Supermarket] SET ARITHABORT OFF 
GO
ALTER DATABASE [Supermarket] SET AUTO_CLOSE ON 
GO
ALTER DATABASE [Supermarket] SET AUTO_SHRINK OFF 
GO
ALTER DATABASE [Supermarket] SET AUTO_UPDATE_STATISTICS ON 
GO
ALTER DATABASE [Supermarket] SET CURSOR_CLOSE_ON_COMMIT OFF 
GO
ALTER DATABASE [Supermarket] SET CURSOR_DEFAULT  GLOBAL 
GO
ALTER DATABASE [Supermarket] SET CONCAT_NULL_YIELDS_NULL OFF 
GO
ALTER DATABASE [Supermarket] SET NUMERIC_ROUNDABORT OFF 
GO
ALTER DATABASE [Supermarket] SET QUOTED_IDENTIFIER OFF 
GO
ALTER DATABASE [Supermarket] SET RECURSIVE_TRIGGERS OFF 
GO
ALTER DATABASE [Supermarket] SET  DISABLE_BROKER 
GO
ALTER DATABASE [Supermarket] SET AUTO_UPDATE_STATISTICS_ASYNC OFF 
GO
ALTER DATABASE [Supermarket] SET DATE_CORRELATION_OPTIMIZATION OFF 
GO
ALTER DATABASE [Supermarket] SET TRUSTWORTHY OFF 
GO
ALTER DATABASE [Supermarket] SET ALLOW_SNAPSHOT_ISOLATION OFF 
GO
ALTER DATABASE [Supermarket] SET PARAMETERIZATION SIMPLE 
GO
ALTER DATABASE [Supermarket] SET READ_COMMITTED_SNAPSHOT OFF 
GO
ALTER DATABASE [Supermarket] SET HONOR_BROKER_PRIORITY OFF 
GO
ALTER DATABASE [Supermarket] SET RECOVERY SIMPLE 
GO
ALTER DATABASE [Supermarket] SET  MULTI_USER 
GO
ALTER DATABASE [Supermarket] SET PAGE_VERIFY CHECKSUM  
GO
ALTER DATABASE [Supermarket] SET DB_CHAINING OFF 
GO
ALTER DATABASE [Supermarket] SET FILESTREAM( NON_TRANSACTED_ACCESS = OFF ) 
GO
ALTER DATABASE [Supermarket] SET TARGET_RECOVERY_TIME = 60 SECONDS 
GO
ALTER DATABASE [Supermarket] SET DELAYED_DURABILITY = DISABLED 
GO
ALTER DATABASE [Supermarket] SET QUERY_STORE = OFF
GO
USE [Supermarket]
GO
/****** Object:  UserDefinedFunction [dbo].[UF_GetBalance]    Script Date: 10/7/2020 5:34:59 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
create function [dbo].[UF_GetBalance](@Product char(3),@Warehouse int) 
returns int
begin
	declare @Balance int
	select @Balance = sum(Change) from dbo.StockHistory
		where ProductCode = @Product 
		and InWarehouse = @Warehouse
	return @Balance
end
GO
/****** Object:  Table [dbo].[StockHistory]    Script Date: 10/7/2020 5:34:59 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[StockHistory](
	[ProductCode] [char](3) NOT NULL,
	[InWarehouse] [int] NOT NULL,
	[Day] [date] NOT NULL,
	[Time] [time](7) NOT NULL,
	[Change] [int] NULL,
	[DoneBy] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[ProductCode] ASC,
	[InWarehouse] ASC,
	[Day] ASC,
	[Time] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  UserDefinedFunction [dbo].[UF_GetStockHistory]    Script Date: 10/7/2020 5:35:00 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
create function [dbo].[UF_GetStockHistory](@StartDay date,@EndDay date,@Product char(3),@Warehouse int)
returns table
as return
	select * from dbo.StockHistory
		where Day >= @StartDay and Day <= @EndDay
		and ProductCode = @Product
		and InWarehouse = @Warehouse
GO
/****** Object:  Table [dbo].[Customer]    Script Date: 10/7/2020 5:35:00 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Customer](
	[Number] [int] NOT NULL,
	[Name] [nvarchar](50) NULL,
	[Year] [int] NULL,
	[Nationality] [nvarchar](20) NULL,
	[Male] [bit] NULL
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Employee]    Script Date: 10/7/2020 5:35:00 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Employee](
	[ID] [int] NOT NULL,
	[FullName] [nvarchar](40) NOT NULL,
	[Male] [bit] NULL,
	[Nationality] [text] NULL,
	[DirectManager] [int] NULL,
	[Password] [varchar](20) NULL,
PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Product]    Script Date: 10/7/2020 5:35:00 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Product](
	[Code] [char](3) NOT NULL,
	[Name] [nvarchar](50) NOT NULL,
	[Type] [char](3) NULL,
	[Price] [int] NULL,
 CONSTRAINT [PK__Product__A25C5AA6D7C9FE4E] PRIMARY KEY CLUSTERED 
(
	[Code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[ProductLocation]    Script Date: 10/7/2020 5:35:00 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[ProductLocation](
	[Floor] [int] NOT NULL,
	[Stall] [char](3) NOT NULL,
	[ContainProduct] [char](3) NULL,
PRIMARY KEY CLUSTERED 
(
	[Floor] ASC,
	[Stall] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[ProductType]    Script Date: 10/7/2020 5:35:00 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[ProductType](
	[Code] [char](3) NOT NULL,
	[Name] [varchar](20) NOT NULL,
	[ManagedBy] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[Code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Timesheet]    Script Date: 10/7/2020 5:35:00 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Timesheet](
	[Employee] [int] NOT NULL,
	[WorkOnFloor] [int] NOT NULL,
	[WorkAtStall] [char](3) NOT NULL,
	[NightShift] [bit] NULL,
PRIMARY KEY CLUSTERED 
(
	[Employee] ASC,
	[WorkOnFloor] ASC,
	[WorkAtStall] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[USDVND]    Script Date: 10/7/2020 5:35:00 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[USDVND](
	[Time] [datetime] NOT NULL,
	[Exchange] [int] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[Time] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Version]    Script Date: 10/7/2020 5:35:00 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Version](
	[Time] [datetime] NULL,
	[Version] [varchar](10) NOT NULL,
 CONSTRAINT [PK_VERSION] PRIMARY KEY CLUSTERED 
(
	[Version] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[WareHouse]    Script Date: 10/7/2020 5:35:00 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[WareHouse](
	[Number] [int] NOT NULL,
	[District] [nvarchar](10) NULL,
	[City] [nvarchar](10) NOT NULL,
	[ManagedBy] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[Number] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [NonClusteredIndex-Type]    Script Date: 10/7/2020 5:35:00 PM ******/
CREATE NONCLUSTERED INDEX [NonClusteredIndex-Type] ON [dbo].[Product]
(
	[Type] ASC
)
INCLUDE([Name]) WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, DROP_EXISTING = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO
ALTER TABLE [dbo].[Employee]  WITH CHECK ADD FOREIGN KEY([DirectManager])
REFERENCES [dbo].[Employee] ([ID])
GO
ALTER TABLE [dbo].[Product]  WITH CHECK ADD  CONSTRAINT [FK__Product__Type__59063A47] FOREIGN KEY([Type])
REFERENCES [dbo].[ProductType] ([Code])
GO
ALTER TABLE [dbo].[Product] CHECK CONSTRAINT [FK__Product__Type__59063A47]
GO
ALTER TABLE [dbo].[ProductLocation]  WITH CHECK ADD  CONSTRAINT [FK__ProductLo__Conta__5EBF139D] FOREIGN KEY([ContainProduct])
REFERENCES [dbo].[Product] ([Code])
GO
ALTER TABLE [dbo].[ProductLocation] CHECK CONSTRAINT [FK__ProductLo__Conta__5EBF139D]
GO
ALTER TABLE [dbo].[ProductType]  WITH CHECK ADD  CONSTRAINT [FK_ProductType_ManagedBy_Employee] FOREIGN KEY([ManagedBy])
REFERENCES [dbo].[Employee] ([ID])
GO
ALTER TABLE [dbo].[ProductType] CHECK CONSTRAINT [FK_ProductType_ManagedBy_Employee]
GO
ALTER TABLE [dbo].[StockHistory]  WITH CHECK ADD  CONSTRAINT [FK_StockHistory_DoneBy_Employee] FOREIGN KEY([DoneBy])
REFERENCES [dbo].[Employee] ([ID])
GO
ALTER TABLE [dbo].[StockHistory] CHECK CONSTRAINT [FK_StockHistory_DoneBy_Employee]
GO
ALTER TABLE [dbo].[StockHistory]  WITH CHECK ADD  CONSTRAINT [FK_StockHistory_InWarehouse_Warehouse] FOREIGN KEY([InWarehouse])
REFERENCES [dbo].[WareHouse] ([Number])
GO
ALTER TABLE [dbo].[StockHistory] CHECK CONSTRAINT [FK_StockHistory_InWarehouse_Warehouse]
GO
ALTER TABLE [dbo].[StockHistory]  WITH CHECK ADD  CONSTRAINT [FK_StockHistory_ProductCode_Product] FOREIGN KEY([ProductCode])
REFERENCES [dbo].[Product] ([Code])
GO
ALTER TABLE [dbo].[StockHistory] CHECK CONSTRAINT [FK_StockHistory_ProductCode_Product]
GO
ALTER TABLE [dbo].[Timesheet]  WITH CHECK ADD FOREIGN KEY([WorkOnFloor], [WorkAtStall])
REFERENCES [dbo].[ProductLocation] ([Floor], [Stall])
GO
ALTER TABLE [dbo].[Timesheet]  WITH CHECK ADD FOREIGN KEY([Employee])
REFERENCES [dbo].[Employee] ([ID])
GO
ALTER TABLE [dbo].[WareHouse]  WITH CHECK ADD  CONSTRAINT [FK_ManagedBy_Employee] FOREIGN KEY([ManagedBy])
REFERENCES [dbo].[Employee] ([ID])
GO
ALTER TABLE [dbo].[WareHouse] CHECK CONSTRAINT [FK_ManagedBy_Employee]
GO
ALTER TABLE [dbo].[Product]  WITH CHECK ADD CHECK  (([Price]>(0)))
GO
ALTER TABLE [dbo].[StockHistory]  WITH CHECK ADD  CONSTRAINT [CK_Change_StockHistory] CHECK  (([Change]<>(0)))
GO
ALTER TABLE [dbo].[StockHistory] CHECK CONSTRAINT [CK_Change_StockHistory]
GO
ALTER TABLE [dbo].[USDVND]  WITH CHECK ADD CHECK  (([Exchange]>(0)))
GO
USE [master]
GO
ALTER DATABASE [Supermarket] SET  READ_WRITE 
GO
