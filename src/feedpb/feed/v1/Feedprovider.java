// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: feedprovider.proto

package feed.v1;

public final class Feedprovider {
  private Feedprovider() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface ProviderGetRequestOrBuilder extends
      // @@protoc_insertion_point(interface_extends:feed.v1.ProviderGetRequest)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
     */
    boolean hasRequest();
    /**
     * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
     */
    feed.v1.Common.FeedGetRequest getRequest();
    /**
     * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
     */
    feed.v1.Common.FeedGetRequestOrBuilder getRequestOrBuilder();
  }
  /**
   * Protobuf type {@code feed.v1.ProviderGetRequest}
   */
  public  static final class ProviderGetRequest extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:feed.v1.ProviderGetRequest)
      ProviderGetRequestOrBuilder {
    // Use ProviderGetRequest.newBuilder() to construct.
    private ProviderGetRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private ProviderGetRequest() {
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return com.google.protobuf.UnknownFieldSet.getDefaultInstance();
    }
    private ProviderGetRequest(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      this();
      int mutable_bitField0_ = 0;
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            default: {
              if (!input.skipField(tag)) {
                done = true;
              }
              break;
            }
            case 10: {
              feed.v1.Common.FeedGetRequest.Builder subBuilder = null;
              if (request_ != null) {
                subBuilder = request_.toBuilder();
              }
              request_ = input.readMessage(feed.v1.Common.FeedGetRequest.parser(), extensionRegistry);
              if (subBuilder != null) {
                subBuilder.mergeFrom(request_);
                request_ = subBuilder.buildPartial();
              }

              break;
            }
          }
        }
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(this);
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(
            e).setUnfinishedMessage(this);
      } finally {
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return feed.v1.Feedprovider.internal_static_feed_v1_ProviderGetRequest_descriptor;
    }

    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return feed.v1.Feedprovider.internal_static_feed_v1_ProviderGetRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              feed.v1.Feedprovider.ProviderGetRequest.class, feed.v1.Feedprovider.ProviderGetRequest.Builder.class);
    }

    public static final int REQUEST_FIELD_NUMBER = 1;
    private feed.v1.Common.FeedGetRequest request_;
    /**
     * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
     */
    public boolean hasRequest() {
      return request_ != null;
    }
    /**
     * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
     */
    public feed.v1.Common.FeedGetRequest getRequest() {
      return request_ == null ? feed.v1.Common.FeedGetRequest.getDefaultInstance() : request_;
    }
    /**
     * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
     */
    public feed.v1.Common.FeedGetRequestOrBuilder getRequestOrBuilder() {
      return getRequest();
    }

    private byte memoizedIsInitialized = -1;
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (request_ != null) {
        output.writeMessage(1, getRequest());
      }
    }

    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (request_ != null) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(1, getRequest());
      }
      memoizedSize = size;
      return size;
    }

    private static final long serialVersionUID = 0L;
    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof feed.v1.Feedprovider.ProviderGetRequest)) {
        return super.equals(obj);
      }
      feed.v1.Feedprovider.ProviderGetRequest other = (feed.v1.Feedprovider.ProviderGetRequest) obj;

      boolean result = true;
      result = result && (hasRequest() == other.hasRequest());
      if (hasRequest()) {
        result = result && getRequest()
            .equals(other.getRequest());
      }
      return result;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptorForType().hashCode();
      if (hasRequest()) {
        hash = (37 * hash) + REQUEST_FIELD_NUMBER;
        hash = (53 * hash) + getRequest().hashCode();
      }
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static feed.v1.Feedprovider.ProviderGetRequest parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static feed.v1.Feedprovider.ProviderGetRequest parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static feed.v1.Feedprovider.ProviderGetRequest parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static feed.v1.Feedprovider.ProviderGetRequest parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static feed.v1.Feedprovider.ProviderGetRequest parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static feed.v1.Feedprovider.ProviderGetRequest parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static feed.v1.Feedprovider.ProviderGetRequest parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static feed.v1.Feedprovider.ProviderGetRequest parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static feed.v1.Feedprovider.ProviderGetRequest parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static feed.v1.Feedprovider.ProviderGetRequest parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(feed.v1.Feedprovider.ProviderGetRequest prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code feed.v1.ProviderGetRequest}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:feed.v1.ProviderGetRequest)
        feed.v1.Feedprovider.ProviderGetRequestOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return feed.v1.Feedprovider.internal_static_feed_v1_ProviderGetRequest_descriptor;
      }

      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return feed.v1.Feedprovider.internal_static_feed_v1_ProviderGetRequest_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                feed.v1.Feedprovider.ProviderGetRequest.class, feed.v1.Feedprovider.ProviderGetRequest.Builder.class);
      }

      // Construct using feed.v1.Feedprovider.ProviderGetRequest.newBuilder()
      private Builder() {
        maybeForceBuilderInitialization();
      }

      private Builder(
          com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
        super(parent);
        maybeForceBuilderInitialization();
      }
      private void maybeForceBuilderInitialization() {
        if (com.google.protobuf.GeneratedMessageV3
                .alwaysUseFieldBuilders) {
        }
      }
      public Builder clear() {
        super.clear();
        if (requestBuilder_ == null) {
          request_ = null;
        } else {
          request_ = null;
          requestBuilder_ = null;
        }
        return this;
      }

      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return feed.v1.Feedprovider.internal_static_feed_v1_ProviderGetRequest_descriptor;
      }

      public feed.v1.Feedprovider.ProviderGetRequest getDefaultInstanceForType() {
        return feed.v1.Feedprovider.ProviderGetRequest.getDefaultInstance();
      }

      public feed.v1.Feedprovider.ProviderGetRequest build() {
        feed.v1.Feedprovider.ProviderGetRequest result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      public feed.v1.Feedprovider.ProviderGetRequest buildPartial() {
        feed.v1.Feedprovider.ProviderGetRequest result = new feed.v1.Feedprovider.ProviderGetRequest(this);
        if (requestBuilder_ == null) {
          result.request_ = request_;
        } else {
          result.request_ = requestBuilder_.build();
        }
        onBuilt();
        return result;
      }

      public Builder clone() {
        return (Builder) super.clone();
      }
      public Builder setField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          Object value) {
        return (Builder) super.setField(field, value);
      }
      public Builder clearField(
          com.google.protobuf.Descriptors.FieldDescriptor field) {
        return (Builder) super.clearField(field);
      }
      public Builder clearOneof(
          com.google.protobuf.Descriptors.OneofDescriptor oneof) {
        return (Builder) super.clearOneof(oneof);
      }
      public Builder setRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          int index, Object value) {
        return (Builder) super.setRepeatedField(field, index, value);
      }
      public Builder addRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          Object value) {
        return (Builder) super.addRepeatedField(field, value);
      }
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof feed.v1.Feedprovider.ProviderGetRequest) {
          return mergeFrom((feed.v1.Feedprovider.ProviderGetRequest)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(feed.v1.Feedprovider.ProviderGetRequest other) {
        if (other == feed.v1.Feedprovider.ProviderGetRequest.getDefaultInstance()) return this;
        if (other.hasRequest()) {
          mergeRequest(other.getRequest());
        }
        onChanged();
        return this;
      }

      public final boolean isInitialized() {
        return true;
      }

      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        feed.v1.Feedprovider.ProviderGetRequest parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (feed.v1.Feedprovider.ProviderGetRequest) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }

      private feed.v1.Common.FeedGetRequest request_ = null;
      private com.google.protobuf.SingleFieldBuilderV3<
          feed.v1.Common.FeedGetRequest, feed.v1.Common.FeedGetRequest.Builder, feed.v1.Common.FeedGetRequestOrBuilder> requestBuilder_;
      /**
       * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
       */
      public boolean hasRequest() {
        return requestBuilder_ != null || request_ != null;
      }
      /**
       * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
       */
      public feed.v1.Common.FeedGetRequest getRequest() {
        if (requestBuilder_ == null) {
          return request_ == null ? feed.v1.Common.FeedGetRequest.getDefaultInstance() : request_;
        } else {
          return requestBuilder_.getMessage();
        }
      }
      /**
       * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
       */
      public Builder setRequest(feed.v1.Common.FeedGetRequest value) {
        if (requestBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          request_ = value;
          onChanged();
        } else {
          requestBuilder_.setMessage(value);
        }

        return this;
      }
      /**
       * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
       */
      public Builder setRequest(
          feed.v1.Common.FeedGetRequest.Builder builderForValue) {
        if (requestBuilder_ == null) {
          request_ = builderForValue.build();
          onChanged();
        } else {
          requestBuilder_.setMessage(builderForValue.build());
        }

        return this;
      }
      /**
       * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
       */
      public Builder mergeRequest(feed.v1.Common.FeedGetRequest value) {
        if (requestBuilder_ == null) {
          if (request_ != null) {
            request_ =
              feed.v1.Common.FeedGetRequest.newBuilder(request_).mergeFrom(value).buildPartial();
          } else {
            request_ = value;
          }
          onChanged();
        } else {
          requestBuilder_.mergeFrom(value);
        }

        return this;
      }
      /**
       * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
       */
      public Builder clearRequest() {
        if (requestBuilder_ == null) {
          request_ = null;
          onChanged();
        } else {
          request_ = null;
          requestBuilder_ = null;
        }

        return this;
      }
      /**
       * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
       */
      public feed.v1.Common.FeedGetRequest.Builder getRequestBuilder() {
        
        onChanged();
        return getRequestFieldBuilder().getBuilder();
      }
      /**
       * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
       */
      public feed.v1.Common.FeedGetRequestOrBuilder getRequestOrBuilder() {
        if (requestBuilder_ != null) {
          return requestBuilder_.getMessageOrBuilder();
        } else {
          return request_ == null ?
              feed.v1.Common.FeedGetRequest.getDefaultInstance() : request_;
        }
      }
      /**
       * <code>optional .feed.v1.FeedGetRequest request = 1;</code>
       */
      private com.google.protobuf.SingleFieldBuilderV3<
          feed.v1.Common.FeedGetRequest, feed.v1.Common.FeedGetRequest.Builder, feed.v1.Common.FeedGetRequestOrBuilder> 
          getRequestFieldBuilder() {
        if (requestBuilder_ == null) {
          requestBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
              feed.v1.Common.FeedGetRequest, feed.v1.Common.FeedGetRequest.Builder, feed.v1.Common.FeedGetRequestOrBuilder>(
                  getRequest(),
                  getParentForChildren(),
                  isClean());
          request_ = null;
        }
        return requestBuilder_;
      }
      public final Builder setUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return this;
      }

      public final Builder mergeUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return this;
      }


      // @@protoc_insertion_point(builder_scope:feed.v1.ProviderGetRequest)
    }

    // @@protoc_insertion_point(class_scope:feed.v1.ProviderGetRequest)
    private static final feed.v1.Feedprovider.ProviderGetRequest DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new feed.v1.Feedprovider.ProviderGetRequest();
    }

    public static feed.v1.Feedprovider.ProviderGetRequest getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<ProviderGetRequest>
        PARSER = new com.google.protobuf.AbstractParser<ProviderGetRequest>() {
      public ProviderGetRequest parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
          return new ProviderGetRequest(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<ProviderGetRequest> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<ProviderGetRequest> getParserForType() {
      return PARSER;
    }

    public feed.v1.Feedprovider.ProviderGetRequest getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_feed_v1_ProviderGetRequest_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_feed_v1_ProviderGetRequest_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\022feedprovider.proto\022\007feed.v1\032\014common.pr" +
      "oto\">\n\022ProviderGetRequest\022(\n\007request\030\001 \001" +
      "(\0132\027.feed.v1.FeedGetRequest2L\n\014FeedProvi" +
      "der\022<\n\003Get\022\033.feed.v1.ProviderGetRequest\032" +
      "\030.feed.v1.FeedGetResponseB\010Z\006feedpbb\006pro" +
      "to3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          feed.v1.Common.getDescriptor(),
        }, assigner);
    internal_static_feed_v1_ProviderGetRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_feed_v1_ProviderGetRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_feed_v1_ProviderGetRequest_descriptor,
        new java.lang.String[] { "Request", });
    feed.v1.Common.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
